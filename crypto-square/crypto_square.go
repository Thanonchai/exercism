package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

var reg *regexp.Regexp

func init() {
	r, err := regexp.Compile("[^a-z0-9]")
	if err != nil {
		panic("Error compiling regex expression")
	}
	reg = r
}

func Encode(pt string) string {
	if pt == "" {
		return ""
	}
	pt = clean(pt)
	row, col := calcRowCol(pt)
	ss := splitRectangle(row, col, pt)
	return strings.Join(ss, " ")
}

func splitRectangle(row, col int, pt string) []string {
	var ss []strings.Builder
	for i := 0; i < col; i++ {
		ss = append(ss, strings.Builder{})
	}
	for i, r := range pt {
		ss[i%col].WriteRune(r)
	}
	result := []string{}
	for i := range ss {
		if ss[i].Len() < row {
			ss[i].WriteString(strings.Repeat(" ", row-ss[i].Len()))
		}
		result = append(result, ss[i].String())
	}
	return result
}

func clean(input string) (output string) {
	input = strings.ToLower(input)
	output = reg.ReplaceAllString(input, "")
	return output
}

func calcRowCol(input string) (row int, col int) {
	sq := math.Sqrt(float64(len(input)))
	row = int(math.Round(sq))
	col = int(math.Ceil(sq))
	return row, col
}
