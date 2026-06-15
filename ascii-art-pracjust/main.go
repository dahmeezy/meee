package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	input := os.Args[1]

	filename := ""

	ban := "standard"

	usage := "Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard"

	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println(usage)
		return
	} else if len(os.Args) == 3 {
		input = os.Args[2]
		if strings.HasPrefix(os.Args[1], "--output=") {
			filename = strings.TrimPrefix(os.Args[1], "--output=")
		} else {
			fmt.Println(usage)
			return
		}
	} else if len(os.Args) == 4 {
		input = os.Args[2]
		ban = os.Args[3]
		if strings.HasPrefix(os.Args[1], "--output=") {
			filename = strings.TrimPrefix(os.Args[1], "--output=")
		} else {
			fmt.Println(usage)
			return
		}

	}

	lines := strings.Split(input, "\\n")

	data, err := os.ReadFile(ban + ".txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	banner := string(data)

	arts := strings.Split(banner, "\n")
	var draw strings.Builder
	for _, text := range lines {
		for row := 0; row < 8; row++ {
			for i := 0; i < len(text); i++ {
				start := (int(text[i])-32)*9 + 1
				result := arts[start+row]
				// fmt.Print(result)
				draw.WriteString(result)
			}
			// fmt.Println()
			draw.WriteString("\n")
		}
		
		
	}


	if filename != "" {
		os.WriteFile(filename, []byte(draw.String()), 0644)
	} else{
		fmt.Println(draw.String())
	}
}
