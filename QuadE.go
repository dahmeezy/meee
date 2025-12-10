package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {

			// Top-left & bottom-left corners
			if (i == 1 && j == 1) || (i == y && j == x) {
				z01.PrintRune('A')

				// Top-right & bottom-right corners
			} else if (i == 1 && j == x) || (i == y && j == 1) {
				z01.PrintRune('C')

				// Borders
			} else if i == 1 || i == y || j == 1 || j == x {
				z01.PrintRune('B')

				// Inside
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
