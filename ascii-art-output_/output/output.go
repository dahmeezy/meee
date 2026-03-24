package output

import (
	"fmt"
	"os"
	"strings"
)

func Asciiout(str, filename, ban string) {

	// Read the content of the specified banner file (e.g., standard.txt) and store
	var draw strings.Builder
	banner, err := os.ReadFile(ban + ".txt")

	if err != nil {
		fmt.Println("Error reading banner file", err)
	}
	// Split the banner file into individual lines
	lines := strings.Split(string(banner), "\n")

	// Split the input string by literal newline characters (\n)(if it contains it)

	words := strings.Split(str, "\\n")

	for _, word := range words {
		if word == " " || word == "" {
			draw.WriteString("\n")
			continue
		}
		for row := 1; row <= 8; row++ {
			for _, char := range word {

				// Find where the character lives in the file
				index := (int(char) - 32) * 9

				ascii := lines[index+row]

				// Append the character's row to the builder
				draw.WriteString(ascii)

			}
			// After finishing one row of all characters, move to next line
			draw.WriteString("\n")
		}
	}
	// Write the final built ASCII string to the specified file
	os.WriteFile(filename, []byte(draw.String()), 0644)

	// Print an empty line (just for spacing in terminal output)
	fmt.Println()
}
