package main

import (
	"fmt"
	"os"
	"fs/fs"
)

func main() {
	ban := ""
	str := ""
	if len(os.Args) == 1 || len(os.Args) > 3 {

		fmt.Println("Correct usage: go run . [string]")
		fmt.Println("Have a specific banner display in mind? Try: go run . [string] [banner]")
		return
	} 
	if len(os.Args) == 2 {
		str = os.Args[1]
		ban = "standard"
	} else if len(os.Args) == 3 {
		str = os.Args[1]
		ban = os.Args[2]
	}
	fs.Ascii(str, ban)
}
