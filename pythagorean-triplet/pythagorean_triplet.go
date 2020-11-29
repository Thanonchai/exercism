package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	result := []Triplet{}
	for i := min; i <= max-2; i++ {
		for j := i + 1; j <= max-1; j++ {
			for k := j + 1; k <= max; k++ {
				if i*i+j*j == k*k {
					result = append(result, Triplet{i, j, k})
				}
			}
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
