package main

import (
	"color/coloring"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		return
	}

	var str, sub string

	if len(os.Args) == 4 {
		str = os.Args[3]

		sub = os.Args[2]
	} else if len(os.Args) == 2 {
		str = os.Args[1]
	} else {
		str = os.Args[2]
	}

	//gets the colorname from the argument passed
	colorName := strings.TrimPrefix(os.Args[1], "--color=")

	bannerData, err := os.ReadFile("standard.txt")

	if err != nil {
		fmt.Println("Error reading banner file", err)
		return
	}

	coloring.Asciicolor(string(bannerData), str, sub, colorName)

}
