package colorout

import (
	"fmt"
	"os"
	"strings"
)

// Used a switch to get the different color types
func GetColor(color string) string {
	color = strings.ToLower(color)
	switch color {
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "purple":
		return "\033[35m"
	case "white":
		return "\033[36m"
	case "orange":
		return "\033[38;5;208m"
	default:
		return "\033[0m"
	}
}

// reading the banner files
func BannerReader(file string) ([]string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading banner: ", err)
	}
	banner := strings.Split(string(data), "\n")
	return banner, err
}

// mark substring/string to be colored in the given string
func Substring2color(str, sub string) []bool {

	lstr := len(str)

	lsub := len(sub)

	//create a slice of bool of the length of string
	checklist := make([]bool, lstr)

	//checks if substring is empty or none is given and sets the checklist true
	if lsub == 0 || sub == "" {

		for i := range checklist {

			checklist[i] = true

		}
		return checklist
	}

	for i := 0; i < lstr; i++ {

		//check if loop is not out of range and check for substring match in the string
		if i+lsub <= lstr && str[i:i+lsub] == sub {

			//sets all the matching substring to true
			for j := i; j < i+lsub; j++ {

				checklist[j] = true
			}

		}
	}

	return checklist
}

func Asciicolor(bannerData, str, sub, colorName string) string{

	var res strings.Builder

	checklist := Substring2color(str, sub)

	// ANSI escape code for the right color
	colorCode := GetColor(colorName)
	// starting to count mulltiple lines
	globalidx := 0
	banner := strings.Split(string(bannerData), "\n")

	// Split the input string by "\n"
	lines := strings.Split(str, "\\n")

	// Loop through each line of input text
	for _, word := range lines {

		if word == "" {
			// this actually detects that this is an empty line and leaves,that means it auhomatically fixes it
			globalidx += len("\\n")
			res.WriteString("\n")
			continue
		}
		// looping through thr ascii charrcters line by line for the banner file
		for i := 1; i <= 8; i++ {
			for ind, char := range word {
				// the ascii formular
				index := (int(char) - 32) * 9
				// checks if this character should be colored.
				if checklist[globalidx+ind] == true {
					res.WriteString(colorCode + banner[index+i] + "\033[0m")

				} else {

					//print plainly
					res.WriteString(banner[index+i])
				}
			}

			res.WriteString("\n")
		}
		// this is actually working for only the len of the word and making the new line only 1 character
		// but the terminal is counting the "\n" as 2 charcters
		//globalidx += len(word) + 1
		// we can actually hardcode it and  just add + 2, but it better to just tell the globalidx thay there is a "\n " and it should add it to the length of words.
		globalidx += len(word) + len("\\n")
	}

	return res.String()
}
