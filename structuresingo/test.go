package main

import "fmt"

type student struct {
	name string
	age  int
}

func changeName(pointer *student, namechosen string) {
	pointer.name = namechosen

}

func main() {
	chris := student{}

	fmt.Println(chris)

	chris.name = "chris"

	chris.age = 30

	fmt.Println(chris)

	changeName(&chris, "Lee")

	fmt.Println(chris)
}

