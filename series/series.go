package series

func All(n int, s string) (out []string) {
	out = make([]string, len(s)-n+1)
	for i := 0; i <= len(s)-n; i++ {
		out[i] = s[i : i+n]
	}
	return out
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return s, false
	}
	return s[0:n], true
}
