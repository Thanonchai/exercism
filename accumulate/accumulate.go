package accumulate

func Accumulate(input []string, op func(string) string) []string {
	output := make([]string, len(input))
	for i, item := range input {
		output[i] = op(item)
	}
	return output
}
