package tools

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
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

// 以json格式输出struct对象
func JSONString(obj interface{}) string {
	o, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(o)
}

func JSONStringPrintln(obj interface{}) string {

	r, err := json.Marshal(obj)
	if err != nil {
		panic(err)

	}
	fmt.Println(string(r))
	return string(r)

}

func StrContains(src []string, tag string) bool {
	for _, s := range src {
		if s == tag {
			return true
		}
	}
	return false
}
