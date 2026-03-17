package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	for _, r := range Itoa(n) {
		z01.PrintRune(r)
	}
}

func Itoa(n int) string {
	result := ""
	if n == 0 {
		return "0"
	}
	if n < 0 {
		n = -n
		result += "-"
	}
	digits := []byte{}
	for n > 0 {
		d := n % 10
		d = '0' + d
		digits = append(digits, byte(d))
		n = n / 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		result += string(digits[i])
	}
	return result
}
