package house

import "fmt"

var actors [][]string = [][]string{
	{"house", "Jack built"},
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
}

const template = "This is%s."

func Verse(v int) string {
	return fmt.Sprintf(template, action(v-1))
}

func action(i int) string {
	var separator string
	if i < 0 {
		return ""
	}
	if i > 0 {
		separator = "\n"
	} else {
		separator = " "
	}
	return fmt.Sprintf(" the %s%sthat %s%s",
		actors[i][0],
		separator,
		actors[i][1],
		action(i-1),
	)
}

func Song() (song string) {
	song = Verse(1)
	for i := 2; i <= 12; i++ {
		song = fmt.Sprintf("%s\n\n%s", song, Verse(i))
	}
	return song
}
