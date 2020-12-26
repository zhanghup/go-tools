package txorm

func (this *Engine) TemplateFuncAdd(name string, f interface{}) {
	this.tmpsync.Lock()
	this.tmps["sql_with_"+name] = f
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
