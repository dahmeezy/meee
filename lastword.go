package piscine

import "github.com/01-edu/z01"

func LastWord(s string) {
	r := []rune(s)
	end := len(r) - 1
	for end >= 0 && (r[end] == ' ' || r[end] == '\t') {
		end--
	}
	if end < 0 {
		z01.PrintRune('\n')
	}
	start := end
	for start >= 0 && (r[start] != ' ' && r[start] != '\t') {
		start--
	}
	for i := start; i < end; i++ {
		z01.PrintRune(r[i])
	}
	z01.PrintRune('\n')
}
