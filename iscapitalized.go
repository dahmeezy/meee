package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	// Check first character
	if s[0] >= 'a' && s[0] <= 'z' {
		return false
	}

	// Check characters after spaces
	for i := 1; i < len(s); i++ {
		if s[i-1] == ' ' {
			if s[i] >= 'a' && s[i] <= 'z' { // lowercase after space = bad
				return false
			}
		}
	}

	return true
}
