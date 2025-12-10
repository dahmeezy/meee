package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.CamelToSnakeCase("HelloWorld"))
	fmt.Println(piscine.CamelToSnakeCase("camelCase"))
	fmt.Println(piscine.CamelToSnakeCase("Already_snake"))
}
