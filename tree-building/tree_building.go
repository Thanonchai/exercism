package tree

import (
	"errors"
)

type Node struct {
	ID       int
	Children []*Node
}

type Record struct {
	ID, Parent int
}

//const root Node = Node{} **Go doesn't support struct constants

func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	indices := make([]int, len(records))
	for i, r := range records {
		if r.ID < 0 || r.ID >= len(records) {
			return nil, errors.New("Invalid record")
		}
		indices[r.ID] = i
	}

	nodes := make([]Node, len(records))
	for i, l := range indices {
		r := records[l]
		if r.ID != i {
			return nil, errors.New("")
		}
		if r.ID > 0 && r.Parent >= r.ID {
			return nil, errors.New("Malformed")
		}
		if r.ID == 0 {
			if r.Parent != 0 {
				return nil, errors.New("Root with Parent")
			} else {
				continue
			}
		}

		nodes[i].ID = i
		p := &nodes[r.Parent]
		p.Children = append(p.Children, &nodes[i])
	}
	return &nodes[0], nil

}
