package summultiples

type void struct{}

var occupied void

func SumMultiples(limit int, divisors ...int) (sum int) {
	if len(divisors) == 0 {
		return 0
	}
	set := map[int]void{}
	for _, num := range divisors {
		if num == 0 {
			continue
		}
		for i := num; i < limit; i += num {
			if _, found := set[i]; !found {
				sum += i
				set[i] = occupied
			}
		}
	}
	return sum
}
