package piscine

import "github.com/01-edu/z01"

func RepeatAlpha(s string) {
	for _, c := range s {
		var count int
		if c >= 'a' && c <= 'z' {
			count = int(c - 'a' + 1)
		} else if c >= 'A' && c <= 'Z' {
			count = int(c - 'A' + 1)
		} else {
			continue // skip non-letters
		}
		for i := 0; i < count; i++ {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
