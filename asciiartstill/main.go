package main

import (
	"colorout/colorout"
	"fmt"
	"os"
	"strings"
)

func main() {

	// controls the length of argumant to be passed
	if len(os.Args) < 2 || len(os.Args) > 6 {
		fmt.Println("Usage: go run . [OUTPUT-FILE-NAME] [STRING] [BANNER] \nYou can use: go run . [OUTPUT-FILE-NAME] [COLOR] [SUBSTRING] [STRING] [BANNER] to get your output in desired color")
		return
	}

	// variable to store expected arguments
	var filename, sub, str, color string
	ban, _ := colorout.BannerReader("standard.txt")

	if len(os.Args) == 2 {

		str = os.Args[1]

	} else if len(os.Args) == 3 {

		// color and str
		if strings.HasPrefix(os.Args[1], "--color=") || strings.HasPrefix(os.Args[1], "--colour=") {
			color = strings.TrimPrefix(os.Args[1], "--color=")
			str = os.Args[2]
			// output and string
		} else if strings.HasPrefix(os.Args[1], "--output=") {
			filename = strings.TrimPrefix(os.Args[1], "--output=")
			str = os.Args[2]
			// str and banner
		} else if !strings.HasPrefix(os.Args[1], "--") {
			str = os.Args[1]
			ban, _ = colorout.BannerReader(os.Args[2] + ".txt")

		} else {
			fmt.Println("Usage: [OPTION(output-file-name or color)] [STRING] \nor [STRING] [BANNER]")
		}

	} else if len(os.Args) == 4 {

		// output,color and string
		if strings.HasPrefix(os.Args[1], "--output") && (strings.HasPrefix(os.Args[2], "--color") || strings.HasPrefix(os.Args[2], "--colour")) {
			filename = strings.TrimPrefix(os.Args[1], "--output=")
			color = strings.TrimPrefix(os.Args[2], "--color=")
			str = os.Args[3]
			// output, string and banner
		} else if strings.HasPrefix(os.Args[1], "--output=") && !strings.HasPrefix(os.Args[2], "--") {
			filename = strings.TrimPrefix(os.Args[1], "--output=")
			str = os.Args[2]
			ban, _ = colorout.BannerReader(os.Args[3] + ".txt")
			// color, substring and string
		} else if (strings.HasPrefix(os.Args[1], "--color") || strings.HasPrefix(os.Args[1], "--colour")) && !strings.HasPrefix(os.Args[2], "--") {
			color = strings.TrimPrefix(os.Args[1], "--color=")
			sub = os.Args[2]
			str = os.Args[3]
			// color string and banner
		} else if (strings.HasPrefix(os.Args[1], "--color") || strings.HasPrefix(os.Args[1], "--colour")) && (os.Args[3] == "standard" || os.Args[3] == "shadow" || os.Args[3] == "thinkertoy") {
			str = os.Args[2]
			ban, _ = colorout.BannerReader(os.Args[3] + ".txt")
		} else {
			fmt.Println("Usage: [OPTION(output-file-name)] [OPTION(color)] [STRING] \nor [OPTION(output-file-name)] [STRING] [BANNER] \nor [OPTION(color)] [SUBSTRING] [STRING] \nor [OPTION(color)] [STRING] [BANNER]")
		}

	} else if len(os.Args) == 5 {
		// output, color, substring and string
		if strings.HasPrefix(os.Args[1], "--output=") && (strings.HasPrefix(os.Args[2], "--color") || strings.HasPrefix(os.Args[2], "--colour")) && (os.Args[4] == "standard" || os.Args[4] == "shadow" || os.Args[4] == "thinkertoy") {
			filename = strings.TrimPrefix(os.Args[1], "--output=")
			color = strings.TrimPrefix(os.Args[2], "--color=")
			str = os.Args[3]
			ban, _ = colorout.BannerReader(os.Args[4] + ".txt")
			// output, color, string and banner
		} else if strings.HasPrefix(os.Args[1], "--output=") && (strings.HasPrefix(os.Args[2], "--color") || strings.HasPrefix(os.Args[2], "--colour")) && !strings.HasPrefix(os.Args[3], "--")  {
			filename = strings.TrimPrefix(os.Args[1], "--output=")
			color = strings.TrimPrefix(os.Args[2], "--color=")
			str = os.Args[4]
			sub=os.Args[3]
			// color, substring, string and banner
		} else if (strings.HasPrefix(os.Args[1], "--color") || strings.HasPrefix(os.Args[1], "--colour"))&&!strings.HasPrefix(os.Args[2], "--")&& (os.Args[4] == "standard" || os.Args[4] == "shadow" || os.Args[4] == "thinkertoy"){
			color = strings.TrimPrefix(os.Args[1], "--color=")
			str = os.Args[3]
			sub=os.Args[2]
			ban, _ = colorout.BannerReader(os.Args[4] + ".txt")
		} else{
		fmt.Println("Usage: [OPTION(output-file-name)] [OPTION(color)] [SUBSTRING] [STRING] \nor [OPTION(output-file-name)] [OPTION(color)] [STRING] [BANNER] \nor [OPTION(color)] [SUBSTRING] [STRING] [BANNER]")
		}	
		

	} else {
// output, color, substring, string and banner
		filename = strings.TrimPrefix(os.Args[1], "--output=")
		color = strings.TrimPrefix(os.Args[2], "--color=")
		sub = os.Args[3]
		str = os.Args[4]
		ban, _ = colorout.BannerReader(os.Args[5] + ".txt")

	}

	if filename != "" {
		// 2. Check if it ends with ".txt"
		if !strings.HasSuffix(filename, ".txt") {
			fmt.Println("Your chosen file name must end with '.txt' ")
			return
		}
	}

	done := colorout.Asciicolor(strings.Join(ban, "\n"), str, sub, color)

	if filename == "" {
		fmt.Print(done) // Just Print to the terminal
	} else {
		os.WriteFile(filename, []byte(done), 0644) // write to the given file/ create one if it doesnt exist and write to it.
	}
}
