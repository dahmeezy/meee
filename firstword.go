package piscine

import "github.com/01-edu/z01"

func FirstWord(s string) {
	started := false

	for _, c := range s {
		// skip leading spaces/tabs
		if !started && (c == ' ' || c == '\t') {
			continue
		}

		// once we hit a letter, start collecting
		if c != ' ' && c != '\t' {
			started = true
			z01.PrintRune(c)
			continue
		}

		// if we are already inside the word and we hit a space → stop
		if started {
			break
		}
	}

	z01.PrintRune('\n')
}
