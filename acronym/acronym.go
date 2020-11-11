package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z']+`)
	words := re.Split(s, -1)
	var result strings.Builder
	for _, word := range words {
		r := string([]rune(word)[0])
		result.WriteString(strings.ToUpper(r))
	}
	return result.String()
}
