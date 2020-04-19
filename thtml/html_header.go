package thtml

import (
	"regexp"
)

func (this *myHtml) Charset() string {
	c, ok := this.Header["charset"]
	if ok {
		return c
	}

	charset, ok := this.document.Find("head").Find("meta[http-equiv]").Attr("content")
	if ok {
		if len(charset) > 0 {
			ss := regexp.MustCompile(`charset=(.*)`).FindAllStringSubmatch(charset, -1)
			if len(ss) > 0 && len(ss[0]) > 0 {
				for i, sss := range ss[0] {
					if i == 0 {
						continue
					}
					if len(sss) > 0 {
						charset = sss
					}
				}
			}
		}
	}
	if len(charset) > 0 {
		this.Header["charset"] = charset
		return charset
	}

	charset, ok = this.document.Find("head").Find("meta[charset]").Attr("charset")
	if ok {
		this.Header["charset"] = charset
		return charset
	}
	charset = "utf8"
	this.Header["charset"] = charset
	return charset
}

func (this *myHtml) Title() string {
	title, ok := this.Header["title"]
	if ok {
		return title
	}
	title = this.document.Find("head").Find("title").Text()
	this.Header["title"] = title
	return title
}
