package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(s string) bool {
	s = strings.ToLower(s)

	for i, c := range s {
		if unicode.IsLetter(c) && strings.ContainsRune(s[i+1:], c) {
			return false
		}
	}

	return true
}
