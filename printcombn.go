package piscine

import "github.com/01-edu/z01"

// PrintCombN prints all possible combinations of n different digits in ascending order
func PrintCombN(n int) {
	// Edge case: if n is not in the valid range, return
	if n <= 0 || n >= 10 {
		return
	}

	// Initialize the digits array for the combination
	digits := make([]int, n)

	// Fill the array with the first combination (e.g., [0, 1, 2] for n=3)
	for i := 0; i < n; i++ {
		digits[i] = i
	}

	// Print the first combination
	printCombination(digits, n)

	// Generate the next combinations
	for {
		i := n - 1

		// Find the rightmost digit that can be incremented
		for i >= 0 && digits[i] == 9-(n-1-i) {
			i--
		}

		// If no more combinations are possible, break the loop
		if i < 0 {
			break
		}

		// Increment the digit and reset the subsequent digits
		digits[i]++
		for j := i + 1; j < n; j++ {
			digits[j] = digits[j-1] + 1
		}

		// Print a separator (comma and space) between combinations
		z01.PrintRune(',')
		z01.PrintRune(' ')

		// Print the current combination
		printCombination(digits, n)
	}

	// Print a newline at the end
	z01.PrintRune('\n')
}

// printCombination prints a combination of digits
func printCombination(digits []int, n int) {
	// Loop through the digits array and print each digit
	for i := 0; i < n; i++ {
		z01.PrintRune(rune(digits[i] + '0'))
	}
}
