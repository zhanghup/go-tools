package tools

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"text/template"
	"time"
)

func ObjectString() *string {
	str := bson.NewObjectId().Hex()
	return &str
}
func Password(password, slat string) string {
	sh := sha256.New()
	sh.Write([]byte(password))
	bts := sh.Sum([]byte(slat))
	return fmt.Sprintf("%x", bts)
}
func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

type date struct {
	time time.Time
}

func Date() date {
	return date{time: time.Now()}
}
func Date2(t time.Time) date {
	return date{time: t}
}

func (this date) HMS() string {
	return this.time.Format("15:04:05")
}
func (this date) YMD() string {
	return this.time.Format("2006-01-02")
}
func (this date) YM() string {
	return this.time.Format("2006-01")
}
func (this date) Y() string {
	return this.time.Format("2006")
}

type ptr struct{}

func Ptr() ptr {
	return ptr{}
}
func (ptr) Int(i int) *int {
	return &i
}

func (ptr) String(i string) *string {
	return &i
}

func (ptr) Int64(i int64) *int64 {
	return &i
}

func (ptr) Float64(i float64) *float64 {
	return &i
}

func (ptr) Interface(i interface{}) *interface{} {
	return &i
}

type str struct{}

func Str() str {
	return str{}
}
func (str) Template(str string, format map[string]interface{}, funcMap template.FuncMap) (string, error) {
	tt := template.New(bson.NewObjectId().Hex())
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

// 以json格式输出struct对象
func (str) JSONString(obj interface{}) string {
	o, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(o)
}

func (str) JSONStringPrintln(obj interface{}) string {

	r, err := json.Marshal(obj)
	if err != nil {
		panic(err)

	}
	fmt.Println(string(r))
	return string(r)

}

func (str) StrContains(src []string, tag string) bool {
	for _, s := range src {
		if s == tag {
			return true
		}
	}
	return false
}
