package piscine

// import "github.com/01-edu/z01"

func FirstWord(s string) string {
	started := false
	fword:=""

	for _, c := range s {
		// skip leading spaces/tabs
		if !started && (c == ' ' || c == '\t') {
			continue
		}

		// once we hit a letter, start collecting
		if c != ' ' && c != '\t' {
			started = true
			fword+=string(c)
			continue
		}

		// if we are already inside the word and we hit a space → stop
		if started {
			break
		}
	}
	return fword + "\n"
}
