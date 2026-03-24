package main

import (
	"fmt"
	"os"
	"output/output"
	"strings"
)

func main() {
	var str, filename, ban string

	if len(os.Args) == 1 || len(os.Args) > 4 {

		fmt.Println("Correct simple usage: `go run . [string]`")
		fmt.Println("To print ouput to a specific file in the banner style you want, use: `go run . [OPTION] [STRING] [BANNER]`")
		return

	}

	if len(os.Args) == 2 {
		str = os.Args[1]
		ban = "standard"
		filename = "output.txt"
	} else if len(os.Args) == 3 {
		str = os.Args[2]
		ban = "standard"
		filename = strings.TrimPrefix(os.Args[1], "--output=")
	} else if len(os.Args) == 4 {
		str = os.Args[2]
		ban = os.Args[3]
		filename = strings.TrimPrefix(os.Args[1], "--output=")
	}
	output.Asciiout(str, filename, ban)
}
