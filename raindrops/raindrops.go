package raindrops

import (
	"sort"
	"strconv"
)

var rules = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

func Convert(drop int) string {
	keys := make([]int, len(rules))
	i := 0
	for k := range rules {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	result := ""
	for index := range keys {
		if drop%keys[index] == 0 {
			result += rules[keys[index]]
		}
	}
	if result == "" {
		return strconv.Itoa(drop)
	}
	return result
}
