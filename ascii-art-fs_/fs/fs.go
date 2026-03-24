package fs

import (
	"fmt"
	"os"
	"strings"
)

// Ascii reads a banner file and prints the input string in ASCII art format.
func Asciifs(str string, ban string) {

	// Read the content of the specified banner file (e.g., standard.txt) and store
	banner, err := os.ReadFile(ban + ".txt")
	if err != nil {
		fmt.Println("Error reading banner file")
		return
	}

	// Split the banner file into individual lines
	lines := strings.Split(string(banner), "\n")

	// Split the input string by literal newline characters (\n)(if it contains it)
	words := strings.Split(str, "\\n")

	for _, word := range words {

		if word == "" {
			fmt.Println()
			continue
		}

		for row := 1; row <= 8; row++ {

			for _, char := range word {

				// Find where the character lives in the file
				index := (int(char) - 32) * 9

				fmt.Print(lines[index+row])

			}
			fmt.Println()
		}

	}
}
