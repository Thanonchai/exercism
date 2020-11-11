package sublist

const (
	EQ    = "equal"
	UNE   = "unequal"
	SUB   = "sublist"
	SUPER = "superlist"
)

type Relation string

func Sublist(listOne, listTwo []int) Relation {
	switch {
	case len(listOne) == len(listTwo) && isSub(listOne, listTwo):
		return EQ
	case len(listOne) < len(listTwo) && isSub(listOne, listTwo):
		return SUB
	case len(listOne) > len(listTwo) && isSub(listTwo, listOne):
		return SUPER
	default:
		return UNE
	}
}

// list1 is assumed to be shorter or equal to list2
func isSub(list1, list2 []int) bool {
	if len(list1) == 0 || len(list2) == 0 {
		return true
	}
	//With the above assumption, the least number of iterations required is 1
	//That is, len(list2) - len(list1) is 0, hence needed 1 more in that case
	iterNeeded := len(list2) - len(list1) + 1
	for j := 0; j < iterNeeded; j++ {
		match := true
		for i, n := range list1 {
			if n != list2[j+i] {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}
	return false
}
