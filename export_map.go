package tools

func MapMerge(m1 map[string]any, m2 map[string]any) map[string]any {
	result := map[string]any{}

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
