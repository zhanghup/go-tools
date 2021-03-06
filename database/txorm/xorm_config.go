package txorm

import "context"

func (this *Engine) TemplateFunc(name string, f interface{}) {
	this.tmpsync.Lock()
	this.tmps[name] = f
	this.tmpsync.Unlock()
}

func (this *Engine) TemplateFuncWith(name string, fn func(ctx context.Context) string) {
	this.tmpsync.Lock()
	this.tmpWiths["sql_with_"+name] = fn
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
