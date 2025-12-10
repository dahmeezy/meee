package piscine

func CountChar(str string, c rune) int {
	count := 0
	for _, v := range str {
		if rune(v) == c {
			count++
		}
	}
	return count
}
