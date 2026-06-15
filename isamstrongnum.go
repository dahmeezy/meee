package piscine

import (
	"math"
	"strconv"
)

func Isamstrong(n int) bool {

	var sum int

	a := strconv.Itoa(n)

	l := len(a)

	for _, v := range a {

		digit, _ := strconv.Atoi(string(v))

		cube := math.Pow(float64(digit),float64(l))

		sum += int(cube)

	}

	if n != sum {
		return false
	}
	return true
}
