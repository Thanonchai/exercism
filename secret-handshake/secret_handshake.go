package secret

type handshake struct {
	code uint
	act  string
}

var handshakes []handshake = []handshake{
	{1, "wink"},
	{2, "double blink"},
	{4, "close your eyes"},
	{8, "jump"},
}

func Handshake(code uint) (secret []string) {
	secret = make([]string, 0)

	for _, h := range handshakes {
		if code&h.code == h.code {
			secret = append(secret, h.act)
		}
	}
	if code&16 == 16 {
		return reverse(secret)
	}
	return secret
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
