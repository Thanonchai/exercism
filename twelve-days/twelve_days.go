package twelve

import (
	"fmt"
	"strings"
)

const (
	DAY = iota
	GIFT
)

var xMas = [][]string{
	{"", ""},
	{"first", " a Partridge in a Pear Tree."},
	{"second", " two Turtle Doves, and"},
	{"third", " three French Hens,"},
	{"fourth", " four Calling Birds,"},
	{"fifth", " five Gold Rings,"},
	{"sixth", " six Geese-a-Laying,"},
	{"seventh", " seven Swans-a-Swimming,"},
	{"eighth", " eight Maids-a-Milking,"},
	{"ninth", " nine Ladies Dancing,"},
	{"tenth", " ten Lords-a-Leaping,"},
	{"eleventh", " eleven Pipers Piping,"},
	{"twelfth", " twelve Drummers Drumming,"},
}

func Verse(num int) string {
	if num < 1 || num > 12 {
		return ""
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me:%s", xMas[num][DAY], gift(num))
}

func gift(num int) string {
	if num < 1 || num > 12 {
		return ""
	}
	return fmt.Sprintf("%s%s", xMas[num][GIFT], gift(num-1))
}

func Song() string {
	var b strings.Builder
	for i := 1; i <= 12; i++ {
		fmt.Fprintf(&b, "%s\n", Verse(i))

	}
	return b.String()[:b.Len()-1]
}
