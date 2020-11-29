package pangram

import "strings"

func IsPangram(input string) bool {
	pangram := make([]bool, 26)
	input = strings.ToUpper(input)
	for _, c := range input {
		if c >= 'A' && c <= 'Z' {
			pangram[int(c-'A')] = true
		}
	}
	for _, p := range pangram {
		if !p {
			return false
		}
	}
	return true
}
