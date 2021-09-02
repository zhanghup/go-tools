package extraction

import (
	"github.com/tealeg/xlsx/v3"
	"io"
	"os"
	"strings"
)

type EngineXlsx struct {
	Config *Config
}

func NewEngineXlsx(cfg *Config) IExtraction {
	return &EngineXlsx{Config: cfg}
}

// Open 读取csv文件数据
// filename 文件路径
// cfg 配置文件，可空可接收一个配置参数
func (this *EngineXlsx) Open(filename string) ([]Sheet, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return this.OpenIO(f)
}

// OpenIO 读取csv文件数据
// read 文件流
// cfg 配置文件，可空可接收一个配置参数
func (this *EngineXlsx) OpenIO(read io.Reader) ([]Sheet, error) {
	data, err := io.ReadAll(read)
	if err != nil {
		return nil, err
	}

	info, err := xlsx.OpenBinary(data)
	if err != nil {
		return nil, err
	}

	result := make([]Sheet, 0)

	for _, s := range info.Sheets {
		sheet := NewSheet(this.Config)
		sheet.Name = s.Name

		err := s.ForEachRow(func(row *xlsx.Row) error {

			items := make([]Cell, 0)
			err := row.ForEachCell(func(c *xlsx.Cell) error {
				ty := this.Type(c.Type(), c.NumFmt)
				items = append(items, Cell{
					Config:    this.Config,
					Type:      ty,
					Value:     c.String(),
					Formatter: this.Formatter(ty, c.NumFmt),
				})
				return nil
			})
			if err != nil {
				return err
			}

			sheet.DataList = append(sheet.DataList, items)

			return nil
		})
		if err != nil {
			return nil, err
		}

		result = append(result, *sheet)
	}

	return result, nil

}

func (this *EngineXlsx) Formatter(f CellType, fmt string) string {
	switch f {
	case CellTypeDate:
		return fmt
	default:
		return ""
	}
}

func (this *EngineXlsx) Type(t xlsx.CellType, fmt string) CellType {
	if fmt != "general" {
		s := []string{"ss", "mm", "h", "d", "m", "yyyy", "yy"}
		n := 0
		for _, key := range s {
			if strings.Index(fmt, key) != -1 {
				n += 1
			}
		}
		if n > 2 {
			return CellTypeDate
		}
	}

	switch t {
	case xlsx.CellTypeNumeric:
		return CellTypeFloat
	case xlsx.CellTypeDate:
		return CellTypeDate
	case xlsx.CellTypeBool:
		fallthrough
	case xlsx.CellTypeInline:
		fallthrough
	case xlsx.CellTypeError:
		fallthrough
	case xlsx.CellTypeStringFormula:
		fallthrough
	case xlsx.CellTypeString:
		fallthrough
	default:
		return CellTypeString
	}
}
