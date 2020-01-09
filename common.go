package tools

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"os"
	"strings"
	"text/template"
	"time"
)

func ObjectString() *string {
	str := bson.NewObjectId().Hex()
	return &str
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

func StrContains(src []string, tag string) bool {
	for _, s := range src {
		if s == tag {
			return true
		}
	}
	return false
}
