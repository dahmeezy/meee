package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {

	for i := 0; i < len(arr); i++ {
		left := arr[i] / 16
		right := arr[i] % 16

		if left < 10 {
			z01.PrintRune(rune('0' + left))
		} else {
			z01.PrintRune(rune(('0' + left - 10)))
		}
		if right < 10 {
			z01.PrintRune(rune('0' + right))
		} else {
			z01.PrintRune(rune(('0' + right - 10)))
		}
		z01.PrintRune(' ')

		if (i+1)%4 == 0 || i == len(arr)-1 {
			z01.PrintRune('\n')
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] < 32 || arr[i] > 126 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(arr[i]))
		}
	}
	z01.PrintRune('\n')
}
