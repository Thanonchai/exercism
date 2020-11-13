package flatten

func Flatten(item interface{}) (result []interface{}) {
	result = process(item)
	n := 0
	for _, x := range result {
		if x != nil {
			result[n] = x
			n++
		}
	}
	return result[:n]
}

func process(item interface{}) (result []interface{}) {
	if items, ok := item.([]interface{}); ok {
		for _, v := range items {
			result = append(result, process(v)...)
		}
		return result
	} else {
		return append(result, item)
	}
}
