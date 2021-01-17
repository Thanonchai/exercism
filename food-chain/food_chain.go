package foodchain

import "strings"

type swallow struct {
	animal, exclamation, catch string
}

var foodchain = []swallow{
	{"", "", ""},
	{"fly", "", ""},
	{"spider", "It wriggled and jiggled and tickled inside her.\n", "fly.\n"},
	{"bird", "How absurd to swallow a bird!\n", "spider that wriggled and jiggled and tickled inside her.\n"},
	{"cat", "Imagine that, to swallow a cat!\n", "bird.\n"},
	{"dog", "What a hog, to swallow a dog!\n", "cat.\n"},
	{"goat", "Just opened her throat and swallowed a goat!\n", "dog.\n"},
	{"cow", "I don't know how she swallowed a cow!\n", "goat.\n"},
}

func Verse(v int) string {
	if v == len(foodchain) {
		return "I know an old lady who swallowed a horse.\nShe's dead, of course!"
	}
	var result strings.Builder
	result.WriteString("I know an old lady who swallowed a ")
	result.WriteString(foodchain[v].animal)
	result.WriteString(".\n")
	result.WriteString(foodchain[v].exclamation)
	for i := v; i > 1; i-- {
		result.WriteString("She swallowed the ")
		result.WriteString(foodchain[i].animal)
		result.WriteString(" to catch the ")
		result.WriteString(foodchain[i].catch)
	}
	result.WriteString("I don't know why she swallowed the fly. Perhaps she'll die.")
	return result.String()
}

func Verses(from, to int) string {
	var result strings.Builder
	for i := from; i < to; i++ {
		result.WriteString(Verse(i))
		result.WriteString("\n\n")
	}
	result.WriteString(Verse(to))
	return result.String()
}

func Song() string {
	return Verses(1, 8)
}
