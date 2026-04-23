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

	ban := strings.Split(string(Readban), "\n")

	words := strings.Split(s, " ")
	
	if len(words)==1{
		JustifyLeft(words[0],banner)
		return
	}

	totalWordLength := 0

	for _, word := range words {

		wordLength := 0

		for _, c := range word {
			start := ((int(c) - 32) * 9) + 1
			wordLength += len(ban[start])
		}
		totalWordLength += wordLength
	}
	gap := len(words) - 1

	space := (GetTerminalWidth() - totalWordLength) / gap

	

	for row := 0; row < 8; row++ {
		var res strings.Builder
		for _, word := range words {

			for _, c := range word {

				start := ((int(c) - 32) * 9) + 1
				result := string(ban[start+row])
				res.WriteString(result)

			}
			fmt.Print(res.String())
			fmt.Print(strings.Repeat(" ", space))
		}
		fmt.Println()

	}
	

}

func GetTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	if out, err := cmd.Output(); err == nil {
		parts := strings.Fields(string(out))
		if len(parts) == 2 {
			if w, err := strconv.Atoi(parts[1]); err == nil {
				return w
			}
		}
	}
	return 80
}

