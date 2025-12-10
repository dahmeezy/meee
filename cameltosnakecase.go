package piscine

func CamelToSnakeCase(s string) string {
	result := ""

	for i, c := range s {
		if c >= 'A' && c <= 'Z' {
			if i != 0 {
				result += "_"
			}
			result += string(c + ('a' - 'A'))
		} else {
			result += string(c)
		}
	}

	return result
}
