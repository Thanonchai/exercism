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

	result := -1
	for i := 0; i+span <= len(digits); i++ {
		current := 1
		for j := i; j < i+span; j++ {
			digit := int(digits[j] - '0')
			if digit < 0 || digit > 9 {
				return -1, errors.New("Invalid character found")
			}
			current = current * digit
		}
		if current > result {
			result = current
		}
	}
	return result, nil
}
