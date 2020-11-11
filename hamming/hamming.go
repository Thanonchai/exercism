package hamming

import (
	"errors"
	"unicode/utf8"
)

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("Inequal sequence lengths!")
	}

	count := 0

	for i := range a {
		fromA, _ := utf8.DecodeRuneInString(a[i:])
		fromB, _ := utf8.DecodeRuneInString(b[i:])
		if fromA != fromB {
			count++
		}
	}

	return count, nil
}
