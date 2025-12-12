package piscine

func LastWord(s string) string {
	r := []rune(s)

	// Step 1: skip trailing spaces/tabs
	end := len(r) - 1
	for end >= 0 && (r[end] == ' ' || r[end] == '\t') {
		end--
	}

	// No word found
	if end < 0 {
		return "\n"
	}

	// Step 2: find the start of the last word
	start := end
	for start >= 0 && r[start] != ' ' && r[start] != '\t' {
		start--
	}

	// Step 3: return the substring that is the last word
	return string(r[start+1 : end+1])
}
