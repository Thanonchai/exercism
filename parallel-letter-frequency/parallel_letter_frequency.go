package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap)
	for _, t := range texts {
		go func(txt string) {
			c <- Frequency(txt)
		}(t)
	}
	m := FreqMap{}
	for i := 0; i < len(texts); i++ {
		freq := <-c
		for k, v := range freq {
			m[k] += v
		}
	}
	return m
}
