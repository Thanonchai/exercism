package scrabble

import "strings"

var scores = map[rune]int{}

func init() {
	config := map[string]int{
		"aeioulnrst": 1,
		"dg":         2,
		"bcmp":       3,
		"fhvwy":      4,
		"k":          5,
		"jx":         8,
		"qz":         10,
	}

	for key, value := range config {
		for _, r := range key {
			scores[r] = value
		}
	}
}

func Score(word string) (score int) {
	for _, r := range strings.ToLower(word) {
		score += scores[r]
	}
	return
}
