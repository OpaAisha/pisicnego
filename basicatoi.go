package piscine

func BasicAtoi(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		digit := s[i] - '0' // Convert the character to an integer by subtracting '0'
		result = result*10 + int(digit)
	}

	return result
}
