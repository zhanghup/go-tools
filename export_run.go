package tools

func Run(fn func(), callback ...func(res interface{})) {
	defer func() {
		if r := recover(); r != nil {
			if len(callback) > 0 {
				callback[0](r)
			}
		}
	}()
	fn()
}
