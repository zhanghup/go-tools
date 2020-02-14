package str

import (
	"bytes"
	"fmt"
	"github.com/zhanghup/go-tools/rft"
	"strings"
	"text/template"
)

func S(format string, args ...interface{}) string {
	params := make([]interface{}, 0)
	for _,p := range args {
		params = append(params, rft.RealValue(p))
	}
	return fmt.Sprintf(format, params...)
}

func Template(str string, format map[string]interface{}, funcMap template.FuncMap) (string, error) {
	tt := template.New(Uid())
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
