package luhn

func Valid(num string) bool {
	sum := 0
	digit := 0
	for i := len(num) - 1; i >= 0; i-- {
		c := num[i]
		switch {
		case c == ' ':
			continue
		case c >= '0' && c <= '9':
			value := int(c - '0')
			if digit%2 == 1 {
				value *= 2
			}
			if value > 9 {
				value -= 9
			}
			sum += value
			digit++
		default:
			return false
		}
	}
	return sum%10 == 0 && digit > 1
}
