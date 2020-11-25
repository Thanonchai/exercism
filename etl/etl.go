package etl

import "strings"

func Transform(old map[int][]string) (transformed map[string]int) {
	transformed = make(map[string]int)
	for k, v := range old {
		for _, s := range v {
			transformed[strings.ToLower(s)] = k
		}
	}
	return transformed
}
