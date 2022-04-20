package tgql

import (
	"context"
	"github.com/zhanghup/go-tools"
	"regexp"
	"strings"
)

var sqlFormatRegexp = regexp.MustCompile(`^[0-9a-zA-Z_]+$`)

func (this *Loader) SqlFormat(sqlstr, field string) string {
	if strings.Index(sqlstr, "select") == -1 && sqlFormatRegexp.MatchString(sqlstr) {

		sqlstr = tools.StrTmp(`
			select {{ .table }}.* from {{ .table }} where {{ .table }}.{{ .field }} in :keys
		`, map[string]any{
			"table": sqlstr,
			"field": field,
		}).String()
	}

	return sqlstr
}

func (this *Loader) LoadXormCtxObject(ctx context.Context, sqlstr string, field string, param ...any) IObject {
	return this.LoadXormCtx(ctx, []map[string]any{}, this.SqlFormat(sqlstr, field), func(tempData any) map[string]any {
		data := tempData.([]map[string]any)
		result := map[string]any{}
		for i, o := range data {
			key, ok := o[field].(string)
			if !ok {
				continue
			}
			result[key] = data[i]
		}
		return result
	}, param...)
}

func (this *Loader) LoadXormCtxSlice(ctx context.Context, sqlstr string, field string, param ...any) IObject {

	return this.LoadXormCtx(ctx, []map[string]any{}, this.SqlFormat(sqlstr, field), func(tempData any) map[string]any {
		data := tempData.([]map[string]any)
		tmp := map[string][]map[string]any{}
		for i, o := range data {
			key, ok := o[field].(string)
			if !ok {
				continue
			}
			tmp[key] = append(tmp[key], data[i])
		}

		result := map[string]any{}
		for k, v := range tmp {
			result[k] = v
		}
		return result
	}, param...)
}
