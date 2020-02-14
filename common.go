package common

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync/atomic"
	"text/template"
	"time"
)

func ObjectId() string {
	machineId := func() []byte {
		var sum [3]byte
		id := sum[:]
		hostname, err1 := os.Hostname()
		if err1 != nil {
			n := uint32(time.Now().UnixNano())
			sum[0] = byte(n >> 0)
			sum[1] = byte(n >> 8)
			sum[2] = byte(n >> 16)
			return id
		}
		hw := md5.New()
		hw.Write([]byte(hostname))
		copy(id, hw.Sum(nil))
		return id
	}()
	processId := os.Getpid()
	objectIdCounter := uint32(time.Now().UnixNano())

	var b [12]byte
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))
	// Machine, first 3 bytes of md5(hostname)
	b[4] = machineId[0]
	b[5] = machineId[1]
	b[6] = machineId[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	b[7] = byte(processId >> 8)
	b[8] = byte(processId)
	// Increment, 3 bytes, big endian
	i := atomic.AddUint32(&objectIdCounter, 1)
	b[9] = byte(i >> 16)
	b[10] = byte(i >> 8)
	b[11] = byte(i)
	return hex.EncodeToString(b[:])
}

// 以json格式输出struct对象
func JSONString(obj interface{}) string {
	o, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(o)
}

func Println(obj interface{}) {

	r, err := json.Marshal(obj)
	if err != nil {
		panic(err)

	}
	fmt.Println(string(r))
	os.Stdout.Write(r)
}

func EncryptPassword(password, slat string) string {
	sh := sha256.New()
	sh.Write([]byte(password))
	bts := sh.Sum([]byte(slat))
	return fmt.Sprintf("%x", bts)
}
func EncryptMD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func TimeToHMS() string {
	return time.Now().Format("15:04:05")
}
func TimeToYMD() string {
	return time.Now().Format("2006-01-02")
}
func TimeToYM() string {
	return time.Now().Format("2006-01")
}
func TimeToYear() string {
	return time.Now().Format("2006")
}

func PtrInt(i int) *int {
	return &i
}
func PtrString(i string) *string {
	return &i
}
func PtrInt64(i int64) *int64 {
	return &i
}
func PtrFloat64(i float64) *float64 {
	return &i
}
func PtrInterface(i interface{}) *interface{} {
	return &i
}

func StrTemplate(str string, format map[string]interface{}, funcMap template.FuncMap) (string, error) {
	tt := template.New(ObjectId())
	fmap := template.FuncMap{
		"title": strings.Title,
	}
	if funcMap != nil {
		for k, v := range funcMap {
			fmap[k] = v
		}
	}
	tt, err := tt.Funcs(fmap).Parse(str)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer([]byte{})
	err = tt.Execute(buf, format)
	return buf.String(), err
}
func StrContains(src []string, tag string) bool {
	for _, s := range src {
		if s == tag {
			return true
		}
	}
	return false
}

func RanderString(l int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	res := make([]byte, l)
	for i := 0; i < l; i++ {
		b := r.Intn(26) + 65
		res[i] = byte(b)
	}
	return string(res)
}
