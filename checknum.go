package piscine

func Checknum(s string) bool {
	for _, n := range s {
		if n >= '0' && n <= '9' {
			return true
		}
	}
	return false
}
