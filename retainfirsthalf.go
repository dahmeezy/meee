package piscine

func RetainFirstHalf(str string) string {
	if len(str) < 2 {
		return str
	}
	r := []rune(str)
	half := len(r) / 2
	return (string(r[:half]))
}
