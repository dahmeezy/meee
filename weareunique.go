package piscine

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	count := 0

	for _, r := range str1 {
		if !contains(str2, r) {
			count++
		}
	}

	for _, v := range str2 {
		if !contains(str1, v) {
			count++
		}
	}
	return count
}

func contains(s string, r rune) bool {
	for _, v := range s {
		if r == v {
			return true
		}
	}
	return false
}
