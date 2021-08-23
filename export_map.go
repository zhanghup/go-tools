package tools

func MapMerge(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}

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
