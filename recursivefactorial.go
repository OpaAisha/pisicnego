package piscine

const MaxFactorial = 20 // Factorial above 20 causes overflow in an int

func RecursiveFactorial(nb int) int {
	// Return 0 for negative numbers or numbers above MaxFactorial
	if nb < 0 || nb > MaxFactorial {
		return 0
	}

	// Base case: 0! = 1 and 1! = 1
	if nb == 0 || nb == 1 {
		return 1
	}

	// Recursive case: factorial(n) = n * factorial(n-1)
	result := nb * RecursiveFactorial(nb-1)

	// Check if the result exceeds MaxFactorial
	if result < 0 {
		return 0 // Return 0 if the result overflows
	}

	return result
}
