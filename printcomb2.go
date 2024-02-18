package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for y := '0'; y <= '9'; y++ {
			for x := '0'; x <= '9'; x++ {
				for z := '0'; z <= '9'; z++ {
					if i <= x && y < z {
						z01.PrintRune(i)
						z01.PrintRune(y)
						z01.PrintRune(' ')
						z01.PrintRune(x)
						z01.PrintRune(z)

						if (i != '9' || y != '8') || (x != '9' || z != '9') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					} else if i < x && y >= z {
						z01.PrintRune(i)
						z01.PrintRune(y)
						z01.PrintRune(' ')
						z01.PrintRune(x)
						z01.PrintRune(z)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
