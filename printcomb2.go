package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i1 := '0'; i1 <= '9'; i1++ {
		for i2 := '0'; i2 <= '9'; i2++ {
			for j1 := i1; j1 <= '9'; j1++ {
				for j2 := '0'; j2 <= '9'; j2++ {
					if i1 == j1 && i2 >= j2 {
						continue
					}
					// Print the first digit for first number
					z01.PrintRune(i1)
					z01.PrintRune(i2)
					// Print a space between the two digits
					z01.PrintRune(' ')
					// Print the second digit
					z01.PrintRune(j1)
					z01.PrintRune(j2)
					if !(i1 == '9' && i2 == '8' && j1 == '9' && j2 == '9') {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n') // Newline at the end
}
