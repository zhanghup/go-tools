package txorm

import "context"

func (this *Engine) TemplateFunc(name string, f interface{}) {
	this.tmpsync.Lock()
	this.tmps[name] = f
	this.tmpsync.Unlock()
}

/*
	TemplateFuncWith Sql With 模板，

	初始化模板：
	db.TemplateFunc("users",function(ctx context.Context) string{
		return "select id,name from user"
	})

	select * from {{ sql_name "users" }}
	=>
	{{ sql_with_users }} select * from __sql_with_users
	=>
	with recursive _ as (select 1),__sql_with_users as (select id,name from user) select * from __sql_with_users
*/
func (this *Engine) TemplateFuncWith(name string, fn func(ctx context.Context) string) {
	this.tmpsync.Lock()
	this.tmpWiths["tmp_"+name] = fn
	this.tmpsync.Unlock()
}

func (this *Engine) TemplateFuncCtx(name string, fn func(ctx context.Context) string) {
	this.tmpsync.Lock()
	this.tmpCtxs["ctx_"+name] = fn
	this.tmpsync.Unlock()
}

func (this *Engine) TemplateFuncKeys() []string {
	this.tmpsync.RLock()
	keys := make([]string, len(this.tmps))
	for k := range this.tmps {
		keys = append(keys, k)
	}
	this.tmpsync.RUnlock()
	return keys
}
