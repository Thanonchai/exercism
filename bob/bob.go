package bob

import (
	"strings"
	"unicode"
)

type Remark struct {
	text string
}

func newRemark(text string) Remark {
	return Remark{strings.TrimSpace(text)}
}

func (r Remark) aSilence() bool {
	return r.text == ""
}

func (r Remark) aYelling() bool {
	hasLetter := strings.IndexFunc(r.text, unicode.IsLetter) >= 0
	isUpperCase := strings.ToUpper(r.text) == r.text
	return hasLetter && isUpperCase
}

func (r Remark) aQuestion() bool {
	return strings.HasSuffix(r.text, "?")
}

func (r Remark) aYellingQuestion() bool {
	return r.aYelling() && r.aQuestion()
}

func Hey(remark string) string {
	r := newRemark(remark)
	switch {
	case r.aSilence():
		return "Fine. Be that way!"
	case r.aYellingQuestion():
		return "Calm down, I know what I'm doing!"
	case r.aYelling():
		return "Whoa, chill out!"
	case r.aQuestion():
		return "Sure."
	default:
		return "Whatever."
	}
}
