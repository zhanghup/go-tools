package tgql

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/zhanghup/go-tools"
	"strings"
)

type myfields struct{}

var Field = myfields{}

/*
query S{
	querya{
		a{
			b{
				c{
					d
					e
					f
					g{
						h
						i
						c
					}
				}
			}
		}
	}
}
*/
/*
	root: "a.b.c" // 取a.b.c下面的所有查询的属性，不包括子属性，例如上面的结果为["d","e","f","g"]
	fields: 针对graphql查询中的字段和传给方法的fields数组做一个并集
*/
func (this myfields) Fields(ctx context.Context, root string, fields ...string) []string {
	graphql.GetOperationContext(ctx)
	flist := this.getNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		root,
		0,
	)
	if len(fields) == 0 {
		return flist
	}

	result := make([]string, 0)
	for _, s := range flist {
		if tools.Str.Contains(fields, s) {
			result = append(result, s)
		}
	}

	return result
}
func (this myfields) getNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, root string, idx int) (preloads []string) {
	if root == "" {
		preloads = append(preloads, this.getNestedPreloads2(fields)...)
	} else {
		keys := strings.Split(root, ".")
		for _, column := range fields {
			if column.Name == keys[idx] {
				if len(keys)-1 == idx {
					preloads = append(preloads, this.getNestedPreloads2(graphql.CollectFields(ctx, column.Selections, nil))...)
				} else {
					preloads = append(preloads, this.getNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), column.Name, idx+1)...)
				}
			}
		}
	}

	return
}

func (this myfields) getNestedPreloads2(fields []graphql.CollectedField) (preloads []string) {
	for _, column := range fields {
		preloads = append(preloads, column.Name)
	}
	return
}
