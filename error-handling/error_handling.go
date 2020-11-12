package erratum

import (
	"errors"
)

var trErr TransientError

func Use(o ResourceOpener, input string) (ultErr error) {
	var r Resource
	var e error

	for {
		r, e = o()
		if e != nil {
			if errors.As(e, &trErr) {
				continue
			} else {
				return e
			}
		} else {
			break
		}
	}

	defer r.Close()
	defer func(res Resource) {
		if rec := recover(); rec != nil {
			if frob, ok := rec.(FrobError); ok {
				res.Defrob(frob.defrobTag)
				ultErr = frob
			} else {
				ultErr = rec.(error)
			}
		}
	}(r)
	r.Frob(input)

	return ultErr
}
