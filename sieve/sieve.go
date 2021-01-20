package sieve

func Sieve(limit int) []int {
	r := []int{}
	if limit < 2 {
		return r
	}
	r = append(r, 2)
	if limit == 2 {
		return r
	}
	sieve := make([]bool, (limit-1)/2+1)
	for i := 1; i < len(sieve); i++ {
		if sieve[i] {
			continue
		}
		r = append(r, 2*i+1)
		for j := i; j < len(sieve); j += 2*i + 1 {
			sieve[j] = true
		}
	}

	return r
}
