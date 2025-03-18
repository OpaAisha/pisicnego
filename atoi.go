package piscine

func Atoi(s string) int {
	// Check if the string is empty
	if len(s) == 0 {
		return 0
	}

	result := 0
	sign := 1
	i := 0

	// Handle optional sign at the beginning
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	// Iterate through the string and convert each character to an integer
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0 // Return 0 if the character is not a digit
		}

		// Convert character to integer and build the result
		result = result*10 + int(s[i]-'0')
	}

	// Apply the sign to the result
	return result * sign
}
