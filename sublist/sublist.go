package sublist

import "strings"

const (
	EQ    = "equal"
	UNE   = "unequal"
	SUB   = "sublist"
	SUPER = "superlist"
)

type Relation string

var blank = []int{}

func Sublist(listOne, listTwo []int) Relation {
	if s, ok := blankCases(listOne, listTwo); ok {
		return s
	}

	if s, ok := equal(listOne, listTwo); ok {
		return s
	}

	one := toString(listOne)
	two := toString(listTwo)

	if len(one) < len(two) {
		if strings.Contains(two, one) {
			return SUB
		}
	}

	if len(two) < len(one) {
		if strings.Contains(one, two) {
			return SUPER
		}
	}

	return UNE
}

func blankCases(listOne, listTwo []int) (Relation, bool) {
	if len(listOne) == 0 { //Slices can only be tested with nil
		switch {
		case len(listTwo) == 0:
			return EQ, true
		default:
			return SUB, true
		}
	}
	if len(listTwo) == 0 && len(listOne) > 1 {
		return SUPER, true
	}
	return "", false
}

func equal(listOne, listTwo []int) (Relation, bool) {
	if len(listOne) == len(listTwo) {
		for i := range listOne {
			if listOne[i] != listTwo[i] {
				return UNE, false
			}
		}
		return EQ, true
	}
	return UNE, false
}

func toString(list []int) string {
	r := make([]rune, len(list))
	for i := range list {
		r[i] = '0' + rune(list[i])
	}
	return string(r)
}
