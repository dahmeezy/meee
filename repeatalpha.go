package piscine

func RepeatAlpha(s string) string {
	var result string

	for _, c := range s {
		var count int
		if c >= 'a' && c <= 'z' {
			count = int(c - 'a' + 1)
		} else if c >= 'A' && c <= 'Z' {
			count = int(c - 'A' + 1)
		} else {
			result += string(c)
			continue // skip non-letters
		}
		for i := 0; i < count; i++ {
			result += string(c)
		}
	}
	return result
}
