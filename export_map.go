package tools

// MapMerge 将多个map合并为一个新的map
func MapMerge[Value any](m1 map[string]Value, m2 ...map[string]Value) map[string]Value {
	result := map[string]Value{}

	if m1 != nil {
		for k, v := range m1 {
			result[k] = v
		}
	}

	if len(m2) > 0 {
		for _, mm := range m2 {
			for k, v := range mm {
				result[k] = v
			}
		}
	}
	return result
}
