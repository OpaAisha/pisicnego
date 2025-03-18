package piscine

const MaxInt32 = 2147483647 // Max value for a 32-bit signed integer

func IterativeFactorial(nb int) int {
	// Return 0 for negative numbers or numbers that may cause overflow
	if nb < 0 || nb > MaxInt32 { // 12! is the highest factorial that fits in an int32
		return 0
	}

	result := 1
	for i := 1; i <= nb; i++ {
		result *= i

		// Check for overflow by comparing to MaxInt32
		if result < 0 {
			return 0
		}
	}

	return result
}
