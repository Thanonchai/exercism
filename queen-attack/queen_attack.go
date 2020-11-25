package queenattack

import (
	"errors"
	"math"
)

func CanQueenAttack(w, b string) (bool, error) {
	if w == b {
		return false, errors.New("Same position")
	}
	if !(valid(w) && valid(b)) {
		return false, errors.New("Invalid Input")
	}

	switch {
	case w[0] == b[0]:
		return true, nil
	case w[1] == b[1]:
		return true, nil
	case tangent(w, b) == 1.0:
		return true, nil
	default:
		return false, nil
	}
}

func valid(position string) bool {
	switch {
	case len(position) != 2:
		return false
	case position[0] < 'a' || position[0] > 'h':
		return false
	case position[1] < '1' || position[1] > '8':
		return false
	default:
		return true
	}
}

func tangent(w, b string) float64 {
	wc := float64(w[0])
	wr := float64(w[1])
	bc := float64(b[0])
	br := float64(b[1])
	result := (wr - br) / (wc - bc)
	return math.Abs(result)
}
