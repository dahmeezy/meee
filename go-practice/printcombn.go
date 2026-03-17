package piscine

import "fmt"

// import "github.com/01-edu/z01"

func PrintCombN(n int) {
	omo := make([]int, n)
	for a := 0; a <= n; a++ {
		for i := range omo {
			omo[i] = i
		}
	}
	fmt.Println(omo)
}
