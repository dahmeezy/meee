package piscine

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := ""
	count := 1
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if i+1 < len(runes) && runes[i] == runes[i+1] {
			count++
		} else {
			result += string(rune(count+'0')) + string(runes[i])
			count = 1
		}
	}
	return result
}
