package piscine

func IsPrime(nb int) bool {
	// Handle base cases
	if nb <= 1 {
		return false // 0 and 1 are not prime numbers
	}
	if nb == 2 {
		return true // 2 is prime
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
