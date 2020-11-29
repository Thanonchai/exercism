package pythagorean

type Triplet [3]int

func Range(min, max int) (result []Triplet) {
	for a := min; a <= max-2; a++ {
		b := a + 1
		c := b + 1
		if c*c > a*a+b*b {
			continue
		}
		for c <= max {
			for a*a+b*b > c*c {
				c++
			}
			if a*a+b*b == c*c {
				result = append(result, Triplet{a, b, c})
			}
			b++
		}
	}
	return result
}

func Sum(p int) []Triplet {
	result := []Triplet{}
	for _, t := range Range(1, p) {
		if t[0]+t[1]+t[2] == p {
			result = append(result, t)
		}
	}
	return result
}
