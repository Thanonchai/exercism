// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) (k Kind) {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	sum := a + b + c
	switch {
	case math.IsNaN(sum) || math.IsInf(sum, 0):
		return NaT
	case a == 0.0 || b == 0.0 || c == 0.0:
		return NaT
	case a+b < c || a+c < b || b+c < a:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || b == c || c == a:
		return Iso
	default:
		return Sca
	}
}
