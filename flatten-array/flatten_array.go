package flatten

func Flatten(item interface{}) (result []interface{}) {

	result = []interface{}{}

	switch arr := item.(type) {
	case []interface{}:
		for _, v := range arr {
			result = append(result, Flatten(v)...)
		}
	case interface{}:
		result = append(result, arr)
	}

	return result
}
