package grains

import (
	"errors"
	"math/big"
)

func Square(num int) (uint64, error) {
	if num < 1 || num > 64 {
		return 0, errors.New("Invalid range")
	}
	pow := uint64(num - 1)
	return uint64(1) << pow, nil
}

func Total() uint64 {
	var sum, pow = big.NewInt(2), big.NewInt(64)
	sum.Exp(sum, pow, nil)
	sum.Sub(sum, big.NewInt(1))
	return sum.Uint64()
}
