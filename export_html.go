package tools

import (
	"regexp"
	"strings"
)

type myHtml struct {
	OriginHtml []byte // 原始网页
	Html       []byte // 网页内容
	/*
		charset: gbk/utf-8等字符集编码
	*/
	Header map[string]string
}

func Html(data []byte) *myHtml {
	return &myHtml{
		OriginHtml: data,
		Html:       data,
		Header:     map[string]string{},
	}
}

func (this *myHtml) Charset() string {
	c, ok := this.Header["charset"]
	if ok {
		return c
	}
	r, err := regexp.Compile(`<meta.*?".*?charset=(.*?)".*?>|<meta.*?charset="(.*?)"`)
	if err != nil {
		return ""
	}
	res := r.FindAllStringSubmatch(string(this.Html), -1)
	if len(res) > 0 {
		for i, char := range res[0] {
			if i == 0 {
				continue
			}
			if len(char) > 0 {
				this.Header["charset"] = char
				break
			}
		}
	}
	c, ok = this.Header["charset"]
	if !ok {
		this.Header["charset"] = "utf8"
	}
	return this.Header["charset"]
}

func (this *myHtml) Title() string {
	title, ok := this.Header["title"]
	if ok {
		return title
	}
	r := regexp.MustCompile(`<title>(.*?)</title>`)
	res := r.FindAllStringSubmatch(string(this.Html), -1)
	if len(res) > 0 && len(res[0]) > 0 {
		for i, s := range res[0] {
			if i == 0 {
				continue
			}
			if len(s) > 0 {
				this.Header["title"] = s
			}
		}
	}
	return this.Header["title"]
}

func (this *myHtml) DecodeHtml() error {
	var err error
	switch strings.ToLower(this.Charset()) {
	case "windows-1251":
		this.Html, err = Charset.Windows1251Decode(this.OriginHtml)
	case "gbk":
		this.Html, err = Charset.GBKDecode(this.OriginHtml)
	}
	return err
}

func (this *myHtml) SetCharset(charset string) *myHtml {
	this.Header["charset"] = charset
	return this
}
