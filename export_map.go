package tools

func MapMerge[Value any](m1 map[string]Value, m2 map[string]Value) map[string]Value {
	result := map[string]Value{}

	if m1 != nil {
		for k, v := range m1 {
			result[k] = v
		}
	}

	if m2 != nil {
		for k, v := range m2 {
			result[k] = v
		}
	}

	return result
}
