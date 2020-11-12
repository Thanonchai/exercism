package erratum

func Use(o ResourceOpener, input string) (err error) {
	var r Resource
	for r, err = o(); err != nil; r, err = o() {
		if _, ok := err.(TransientError); ok {
			continue
		} else {
			return err
		}
	}

	defer r.Close()
	defer func() {
		if rec := recover(); rec != nil {
			if frob, ok := rec.(FrobError); ok {
				r.Defrob(frob.defrobTag)
				err = frob
			} else {
				err = rec.(error)
			}
		}
	}()
	r.Frob(input)

	return err
}
