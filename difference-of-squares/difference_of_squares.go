package diffsquares

import "math"

func SquareOfSum(num int) int {
	sum := num * (1 + num) / 2
	return sum * sum
}

func SumOfSquares(num int) int {
	return num * (num + 1) * (2*num + 1) / 6
}

func Difference(num int) int {
	squareOfSum := float64(SquareOfSum(num))
	sumOfSquares := float64(SumOfSquares(num))
	diff := math.Abs(squareOfSum - sumOfSquares)
	return int(diff)
}
