package coloring

import (
	"fmt"
	"strings"
)

func Asciicolor(bannerData, str, sub, colorName string) {

	// boolean array marking which characters should be colored (from your previous Sub2color function).
	checklist := Sub2color(str, sub)

	// ANSI escape code for the chosen color
	colorCode := Getcolor(colorName)

	// keeps track of the position in the full string across multiple lines.
	globalidx := 0

	// Convert file content into lines
	banner := strings.Split(string(bannerData), "\n")

	// Split the input string by "\n"
	lines := strings.Split(str, "\\n")

	// Loop through each line of input text
	for _, word := range lines {

		if word == "" {
			globalidx++
			fmt.Println()
			continue
		}

		// since ach ASCII character is 8 lines tall
		for i := 1; i <= 8; i++ {

			for ind, char := range word {

				// calculates where in the banner the ASCII lines for this character start.
				index := (int(char) - 32) * 9

				// checks if this character should be colored.
				if checklist[globalidx+ind] == true {

					// color it
					fmt.Print(colorCode + banner[index+i] + "\033[0m")

				} else {

					//print plainly
					fmt.Print(banner[index+i])
				}
			}

			fmt.Println()
		}
		globalidx += len(word) + 1
	}
}

func Getcolor(name string) string {

	name = strings.ToLower(name)

	switch name {
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"

	default:
		return "\033[0m"
	}
}
