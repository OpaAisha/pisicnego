package piscine

/*package main*/

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
		return
	}

	// Handle the negative case
	if n < 0 {
		z01.PrintRune('-') // Print the minus sign for negative numbers
		n = -n             // Convert the number to positive for further processing
	}

	// Handle the recursion for printing digits
	if n >= 10 {
		PrintNbr(n / 10) // Recursively print the digits of the number
	}

	// Print the last digit
	z01.PrintRune(rune('0' + n%10)) // Convert the last digit to rune and print it
}
