package main

import (
	"fmt"
	"math"
)

// "fmt"
// "gopractice/justgo"

func main() {
	fmt.Println(FindNb(225))
	fmt.Println(FindNb(1071225))
	fmt.Println(FindNb(91716553919377))
	fmt.Println(FindNb(4183059834009))

	fmt.Println(Palindrome("racecar"))
	// words := []string{"you", "are", "a", "miracle", "a", "blessing", "from", "above"}
	// o:=[]string{"i", "don't", "even","know","me"}

	// fmt.Println(justgo.Capsome(o,2))

}

func FindNb(m int) int {
	s := 8 * math.Sqrt(float64(m))
	n := int((math.Sqrt(s+1) - 1) / 2)
	var sum int
	var count int
	for i := n; i > 0; i-- {
		sum += int(math.Pow(float64(i), 3))
		count++

	}
	if sum == m {
		return count
	} else {
		return -1
	}
}
func Palindrome(s string) bool {
	in := len(s) - 1
	for i := range s {
		if s[i] != s[in-i] {
			return false
		}
	}
	return true
}
