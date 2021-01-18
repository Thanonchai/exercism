package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int, error) {
	switch {
	case span < 0:
		return -1, errors.New("Invalid span.")
	case span == 0:
		return 1, nil
	case span > len(digits):
		return -1, errors.New("Invalid span.")
	}

	s, max, current := 0, 0, 1
	mem := make([]int, span)
	for _, r := range digits {
		d := int(r - '0')
		switch {
		case d < 0 || d > 9:
			return -1, errors.New("Invalid character found")
		case d == 0:
			s = 0
			current = 1
		default:
			current *= d
			if s >= span {
				current /= mem[s%span]
			}
			if s+1 >= span && current > max {
				max = current
			}
			mem[s%span] = d
			s++
		}
	}
	return max, nil
}
