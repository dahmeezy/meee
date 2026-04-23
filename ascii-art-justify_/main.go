package main

import (
	justify "justify/justy"
	"os"
)

func main() {
	str := os.Args[1]

	banner := "standard"

	if len(os.Args)==3{
		banner=os.Args[2]
	}

	justify.Justify(str, banner)

	

}
