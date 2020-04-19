package thtml

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/zhanghup/go-tools"
	"strings"
)

type myHtml struct {
	origin   []byte // 原始网页
	html     []byte // 网页内容
	document *goquery.Document
	/*
		charset: gbk/utf-8等字符集编码
	*/
	Header map[string]string
}

func New(data []byte) (*myHtml, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	return &myHtml{
		origin:   data,
		html:     data,
		document: doc,
		Header:   map[string]string{},
	}, nil
}

// 先确定文件编码，然后解码成utf8
func (this *myHtml) DecodeHtml() error {
	var err error
	switch strings.ToLower(this.Charset()) {
	case "windows-1251":
		this.html, err = tools.Charset.Windows1251Decode(this.origin)
	case "gbk":
		this.html, err = tools.Charset.GBKDecode(this.origin)
	}
	if err != nil {
		return err
	}
	this.document, err = goquery.NewDocumentFromReader(bytes.NewBuffer(this.html))
	return err
}

func (this *myHtml) SetCharset(charset string) *myHtml {
	this.Header["charset"] = charset
	return this
}
