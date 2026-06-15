package justify

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func JustifyLeft(s string, banner string) {
	Readban, _ := os.ReadFile(banner + ".txt")

	ban := strings.Split(string(Readban), "\n")

	lines := strings.Split(s, "\\n")

	for _, w := range lines {

		for row := 0; row <= 8; row++ {

			for _, c := range w {

				start := ((int(c) - 32) * 9) + 1
				result := string(ban[start+row])
				fmt.Print(result)

			}

			fmt.Println()
		}
	}
}

func JustifyRight(s string, banner string) {
	Readban, _ := os.ReadFile(banner + ".txt")

	ban := strings.Split(string(Readban), "\n")

	lines := strings.Split(s, "\\n")

	// Using a slice of 8 strings to represent the 8 rows of ASCII
	r := make([]string, 8)

	for _, w := range lines {

		for row := 0; row < 8; row++ {
			res := ""

			for _, c := range w {

				start := ((int(c) - 32) * 9) + 1
				result := string(ban[start+row])
				res += result

			}

			r[row] = res
		}
	}

	wordLength := len(r[0])

	space := GetTerminalWidth() - wordLength

	for i := 0; i < 8; i++ {
		fmt.Print(strings.Repeat(" ", space))
		fmt.Println(r[i])
	}

}

func JustifyCenter(s string, banner string) {
	Readban, _ := os.ReadFile(banner + ".txt")

	ban := strings.Split(string(Readban), "\n")

	lines := strings.Split(s, "\\n")

	// Using a slice of 8 strings to represent the 8 rows of ASCII
	r := make([]string, 8)

	for _, w := range lines {

		for row := 0; row < 8; row++ {
			res := ""

			for _, c := range w {

				start := ((int(c) - 32) * 9) + 1
				result := string(ban[start+row])
				res += result

			}

			r[row] = res
		}
	}

	wordLength := len(r[0])

	space := (GetTerminalWidth() - wordLength) / 2

	for i := 0; i < 8; i++ {
		fmt.Print(strings.Repeat(" ", space))
		fmt.Println(r[i])
	}
}

func Justify(s string, banner string) {

	Readban, _ := os.ReadFile(banner + ".txt")

	banners := strings.Split(string(Readban), "\n")

	textSlice := strings.Fields(s)

	gap := len(textSlice) - 1

	if gap < 1 {
		gap = 1
	}

	terminalWidth := GetTerminalWidth()

	wordLength := 0

	for row := 0; row < 8; row++ {

		for _, text := range textSlice {
			result := ""
			for col := 0; col < len(text); col++ {
				start := (int(text[col]-32)*9) + 1
				result += banners[start+row]
			}
			if row == 0 {
				wordLength += len(result)

			}
		}

	}
	space := (terminalWidth - wordLength) / gap

	if space < 0{
		fmt.Println("space is too narrow")
		return
	}

	for row := 0; row < 8; row++ {

		for i, text := range textSlice {
			result := ""
			for col := 0; col < len(text); col++ {
				start := int(text[col]-32)*9 + 1
				result += banners[start+row]
			}
			fmt.Print(result)
            if i < len(textSlice) -1{
				fmt.Print(strings.Repeat(" ", space))
			}
			
		}
		fmt.Println()

	}

}

func GetTerminalWidth() int {
	// used to run a command equivalent to typing "stty size" in the terminal
	cmd := exec.Command("stty", "size")
	// connects the command from before to the actual terminal used
	cmd.Stdin = os.Stdin
	// check if we get an output
	if out, err := cmd.Output(); err == nil {
		// split output into slice of strings
		parts := strings.Fields(string(out))
		if len(parts) == 2 {
			// convert and return second value
			if w, err := strconv.Atoi(parts[1]); err == nil {
				return w
			}
		}
	}
	// return this if we are unable to get terminal width
	return 80
}
