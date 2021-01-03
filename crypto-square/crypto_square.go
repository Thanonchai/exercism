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
	result := transpose(ss)
	return strings.Join(result, " ")
}

func transpose(ss []string) []string {
	row := len(ss)
	result := []string{}
	for j := 0; j < len(ss[0]); j++ {
		line := strings.Builder{}
		for k := 0; k < row; k++ {
			line.WriteByte(ss[k][j])
		}
		result = append(result, line.String())
	}
	return result
}

func splitRectangle(row, col int, pt string) []string {
	ss := []string{}
	for i := 0; i+col <= len(pt); i += col {
		ss = append(ss, pt[i:i+col])
	}
	if row*col > len(pt) {
		ss = append(ss, pt[(row-1)*col:]+strings.Repeat(" ", row*col-len(pt)))
	}
	return ss
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
