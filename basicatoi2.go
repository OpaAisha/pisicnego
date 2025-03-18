package piscine

func BasicAtoi2(s string) int {
	result := 0

	// Iterate over each character of the string
	for i := 0; i < len(s); i++ {
		// Check if the current character is a digit
		if s[i] < '0' || s[i] > '9' {
			return 0 // Return 0 if the character is not a digit
		}

		// Convert character to its integer value and build the result
		result = result*10 + int(s[i]-'0')
	}

	return result
}
