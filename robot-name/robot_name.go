package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var (
	names      = make(map[string]bool)
	seededRand *rand.Rand
)

func makeName() string {
	// Rune is an int in utf8-space
	// Adding 0-25 offset by 'A' gives a random letter
	a1 := seededRand.Intn(26) + 'A'
	a2 := seededRand.Intn(26) + 'A'
	d := seededRand.Intn(1000)
	return fmt.Sprintf("%c%c%3d", a1, a2, d)
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(names) >= 26*26*10*10*10 {
		return "", errors.New("Robot names exhausted")
	}
	// Generate until we have a unique nonzero name
	for r.name == "" || names[r.name] == true {
		r.name = makeName()
	}
	names[r.name] = true
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func init() {
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}
