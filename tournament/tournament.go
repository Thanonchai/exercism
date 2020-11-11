package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	Name    string
	W, D, L int
}

var leaders map[string]*Team

func (t *Team) MatchPlay() int {
	return t.W + t.D + t.L
}

func (t *Team) Points() int {
	return 3*t.W + t.D
}

func Tally(r io.Reader, w io.Writer) error {
	leaders = map[string]*Team{}
	reader := bufio.NewScanner(r)
	for reader.Scan() {
		line := reader.Text()
		if len(strings.TrimSpace(line)) == 0 || line[0] == '#' {
			continue
		}
		result := strings.Split(line, ";")

		switch {
		case len(result) != 3:
			return errors.New("A match result should have 3 parts.")
		case result[0] == result[1]:
			return errors.New("Duplicate team names.")
		}

		switch result[2] {
		case "win":
			winner := lookup(result[0])
			loser := lookup(result[1])
			win(winner, loser)
		case "loss":
			winner := lookup(result[1])
			loser := lookup(result[0])
			win(winner, loser)
		case "draw":
			lookup(result[0]).D++
			lookup(result[1]).D++
		default:
			return errors.New("Invalid result")
		}
	}

	w.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	scores := result()
	for _, v := range scores {
		s := fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n",
			v.Name,
			v.MatchPlay(),
			v.W,
			v.D,
			v.L,
			v.Points())
		w.Write([]byte(s))
	}
	return nil
}

func lookup(name string) *Team {
	if obj, ok := leaders[name]; ok {
		return obj
	}
	team := &Team{Name: name}
	leaders[name] = team
	return team
}

func win(t1, t2 *Team) {
	t1.W++
	t2.L++
}

func result() []Team {
	result := []Team{}
	for _, v := range leaders {
		result = append(result, *v)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Points() == result[j].Points() {
			return strings.Compare(result[i].Name, result[j].Name) < 0
		} else {
			return result[i].Points() > result[j].Points()
		}
	})

	return result[:]
}
