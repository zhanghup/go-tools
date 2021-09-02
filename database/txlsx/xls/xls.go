// Package xls implements the Microsoft Excel Binary File Format (.xls) Structure.
// More specifically, it contains just enough detail to extract cell contents,
// data types, and last-calculated formula values. In particular, it does NOT
// implement formatting or formula calculations.
package xls

// https://docs.microsoft.com/en-us/openspecs/office_file_formats/ms-xls/cd03cb5f-ca02-4934-a391-bb674cb8aa06

import (
	"context"
	"encoding/binary"
	"errors"
	"github.com/zhanghup/go-tools/database/txlsx/xls/cfb"
	"github.com/zhanghup/go-tools/database/txlsx/xls/commonxl"
	"github.com/zhanghup/go-tools/database/txlsx/xls/crypto"
	"io"
	"log"
	"sync"
)

// WorkBook represents an Excel workbook containing 1 or more sheets.
type WorkBook struct {
	ctx context.Context
	doc *cfb.Document

	prot     bool
	h        *header
	sheets   []*boundSheet
	codepage uint16
	dateMode uint16
	strings  []string

	password   string
	substreams [][]*rec

	fpos          int64
	pos2substream map[int64]int

	nfmt commonxl.Formatter
	xfs  []uint16
}

func (b *WorkBook) IsProtected() bool {
	return b.prot
}

func OpenIO(read io.Reader) (Source, error) {
	doc, err := cfb.OpenIO(read)
	if err != nil {
		return nil, err
	}

	b := &WorkBook{
		doc: doc,

		pos2substream: make(map[int64]int, 16),
		xfs:           make([]uint16, 0, 128),
	}

	rdr, err := doc.Open("Workbook")
	if err != nil {
		return nil, errors.New("ErrNotInFormat")
	}
	raw, err := io.ReadAll(rdr)
	if err != nil {
		return nil, err
	}

	err = b.loadFromStream(raw)
	return b, err

}

func Open(filename string) (Source, error) {
	doc, err := cfb.Open(filename)
	if err != nil {
		return nil, err
	}

	b := &WorkBook{
		doc: doc,

		pos2substream: make(map[int64]int, 16),
		xfs:           make([]uint16, 0, 128),
	}

	rdr, err := doc.Open("Workbook")
	if err != nil {
		return nil, errors.New("ErrNotInFormat")
	}
	raw, err := io.ReadAll(rdr)
	if err != nil {
		return nil, err
	}

	err = b.loadFromStream(raw)
	return b, err
}

func (b *WorkBook) loadFromStream(raw []byte) error {
	return b.loadFromStream2(raw, false)
}

func (b *WorkBook) loadFromStreamWithDecryptor(raw []byte, dec crypto.Decryptor) error {
	// interestingly (insecurely) BIFF8 keeps Record Types and sizes in the clear,
	// has a few records that are not encrypted, and has 1 record type that does
	// not encrypt the 32bit integer position at the beginning (while encrypting
	// the rest). It also resets the encryption block counter every 1024 bytes
	// (counting all the "skipped" bytes described above).
	//
	// So this code streams the records through the decryption, but also records
	// a set of overlays applied to the final result which restore the
	// "cleartext" contents in line with the decrypted content.

	pos := 0
	zeros := [8224]byte{}

	type overlay struct {
		Pos int

		RecType   recordType
		DataBytes uint16
		Data      []byte // NB len() not necessarily = DataBytes
	}
	replaceBlocks := []overlay{}

	var err error
	for err == nil && len(raw[pos:]) > 4 {
		o := overlay{}
		o.Pos = pos
		o.RecType = recordType(binary.LittleEndian.Uint16(raw[pos : pos+2]))
		o.DataBytes = binary.LittleEndian.Uint16(raw[pos+2 : pos+4])
		pos += 4

		// copy to output and decryption stream
		binary.Write(dec, binary.LittleEndian, o.RecType)
		binary.Write(dec, binary.LittleEndian, o.DataBytes)
		tocopy := int(o.DataBytes)

		switch o.RecType {
		case RecTypeBOF, RecTypeFilePass, RecTypeUsrExcl, RecTypeFileLock, RecTypeInterfaceHdr, RecTypeRRDInfo, RecTypeRRDHead:
			// untouched data goes directly into output
			o.Data = raw[pos : pos+int(o.DataBytes)]
			pos += int(o.DataBytes)
			dec.Write(zeros[:int(o.DataBytes)])
			tocopy = 0

		case RecTypeBoundSheet8:
			// copy 32-bit position to output
			o.Data = raw[pos : pos+4]
			pos += 4
			dec.Write(zeros[:4])
			tocopy -= 4
		}

		if tocopy > 0 {
			_, err = dec.Write(raw[pos : pos+tocopy])
			pos += tocopy
		}
		replaceBlocks = append(replaceBlocks, o)
	}
	dec.Flush()

	alldata := dec.Bytes()
	for _, o := range replaceBlocks {
		offs := int(o.Pos)
		binary.LittleEndian.PutUint16(alldata[offs:], uint16(o.RecType))
		binary.LittleEndian.PutUint16(alldata[offs+2:], uint16(o.DataBytes))
		if len(o.Data) > 0 {
			offs += 4
			copy(alldata[offs:], o.Data)
		}
	}

	// recurse into the stream parser now that things are decrypted
	return b.loadFromStream2(alldata, true)
}

func (b *WorkBook) Close() error {
	// return records to the pool for reuse
	for i, sub := range b.substreams {
		for _, r := range sub {
			r.Data = nil // allow GC
			recPool.Put(r)
		}
		b.substreams[i] = b.substreams[i][:0]
	}
	b.substreams = b.substreams[:0]
	return nil
}

func (b *WorkBook) loadFromStream2(raw []byte, isDecrypted bool) error {
	b.h = &header{}
	substr := -1
	nestedBOF := 0
	b.pos2substream = make(map[int64]int, 10)
	b.fpos = 0

	// IMPORTANT: if there are any existing records, we need to return them to the pool
	for i, sub := range b.substreams {
		for _, r := range sub {
			recPool.Put(r)
		}
		b.substreams[i] = b.substreams[i][:0]
	}
	b.substreams = b.substreams[:0]

	rawfull := raw
	nr, no, err := b.nextRecord(raw)
	for err == nil {
		raw = raw[no:]
		switch nr.RecType {
		case RecTypeEOF:
			nestedBOF--
		case RecTypeBOF:
			// when substreams are nested, keep them in the same grouping
			if nestedBOF == 0 {
				substr = len(b.substreams)
				b.substreams = append(b.substreams, []*rec{})
				b.pos2substream[b.fpos] = substr
			}
			nestedBOF++
		}
		b.fpos += int64(4 + len(nr.Data))

		// if there's a FilePass record, the data is encrypted
		if nr.RecType == RecTypeFilePass && !isDecrypted {
			etype := binary.LittleEndian.Uint16(nr.Data)
			switch etype {
			case 1:
				dec, err := crypto.NewBasicRC4(nr.Data[2:])
				if err != nil {
					log.Println("xls: rc4 encryption failed to set up", err)
					return err
				}
				return b.loadFromStreamWithDecryptor(rawfull, dec)
			case 2, 3, 4:
				log.Println("need Crypto API RC4 decryptor")
				return errors.New("xls: unsupported Crypto API encryption method")
			default:
				return errors.New("xls: unsupported encryption method")
			}
		}

		b.substreams[substr] = append(b.substreams[substr], nr)
		nr, no, err = b.nextRecord(raw)
	}
	if err == io.EOF {
		err = nil
	}
	if err != nil {
		return err
	}

	for _, records := range b.substreams {

		for i, nr := range records {
			if len(nr.Data) == 0 {
				continue
			}

			switch nr.RecType {
			case RecTypeSST:
				// Shared String Table is often continued across multiple records,
				// so we want to gather them all before starting to parse (some
				// strings may span the gap between records)
				recSet := []*rec{nr}

				lastIndex := i
				for len(records) > (lastIndex+1) && records[lastIndex+1].RecType == RecTypeContinue {
					lastIndex++
					recSet = append(recSet, records[lastIndex])
				}

				b.strings, err = parseSST(recSet)
				if err != nil {
					return err
				}

			case RecTypeContinue:
				// no-op (used above)
			case RecTypeEOF:
				// done

			case RecTypeBOF:
				b.h = &header{
					Version:  binary.LittleEndian.Uint16(nr.Data[0:2]),
					DocType:  binary.LittleEndian.Uint16(nr.Data[2:4]),
					RupBuild: binary.LittleEndian.Uint16(nr.Data[4:6]),
					RupYear:  binary.LittleEndian.Uint16(nr.Data[6:8]),
					MiscBits: binary.LittleEndian.Uint64(nr.Data[8:16]),
				}

				if b.h.Version != 0x0600 {
					return errors.New("xls: invalid file version")
				}
				if b.h.RupYear != 0x07CC && b.h.RupYear != 0x07CD {
					return errors.New("xls: unsupported biff version")
				}
				/*
					if b.h.DocType != 0x0005 && b.h.DocType != 0x0010 {
						// we only support the workbook or worksheet substreams
						log.Println("xls: unsupported document type")
						//break
					}
				*/

			case RecTypeCodePage:
				// BIFF8 is entirely UTF-16LE so this is actually ignored
				b.codepage = binary.LittleEndian.Uint16(nr.Data)

			case RecTypeDate1904:
				b.dateMode = binary.LittleEndian.Uint16(nr.Data)

			case RecTypeFormat:
				// Format maps a format ID to a code string
				fmtNo := binary.LittleEndian.Uint16(nr.Data)
				formatStr, _, err := decodeXLUnicodeString(nr.Data[2:])
				if err != nil {
					log.Println("fail2", err)
					return err
				}
				b.nfmt.Add(fmtNo, formatStr)

			case RecTypeXF:
				// XF records merge multiple style and format directives to one ID
				// ignore font id at nr.Data[0:2]
				fmtNo := binary.LittleEndian.Uint16(nr.Data[2:])
				b.xfs = append(b.xfs, fmtNo)

			case RecTypeBoundSheet8:
				// Identifies the postition within the stream, visibility state,
				// and name of a worksheet
				bs := &boundSheet{}
				bs.Position = binary.LittleEndian.Uint32(nr.Data[:4])
				bs.HiddenState = nr.Data[4]
				bs.SheetType = nr.Data[5]

				bs.Name, _, err = decodeShortXLUnicodeString(nr.Data[6:])
				if err != nil {
					return err
				}
				b.sheets = append(b.sheets, bs)
			default:

			}
		}
	}

	return err
}

var recPool = sync.Pool{
	New: func() interface{} {
		return &rec{}
	},
}

func (b *WorkBook) nextRecord(raw []byte) (*rec, int, error) {
	if len(raw) < 4 {
		return nil, 0, io.EOF
	}
	rec := recPool.Get().(*rec)

	rec.RecType = recordType(binary.LittleEndian.Uint16(raw[:2]))
	rec.RecSize = binary.LittleEndian.Uint16(raw[2:4])
	if len(raw[4:]) < int(rec.RecSize) {
		recPool.Put(rec)
		return nil, 4, io.ErrUnexpectedEOF
	}
	rec.Data = raw[4 : 4+rec.RecSize]
	return rec, int(4 + rec.RecSize), nil
}
