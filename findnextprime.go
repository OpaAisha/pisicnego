package piscine

// FindNextPrime returns the first prime number greater than or equal to nb.
func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2 // 2 is the smallest prime number
	}
	// If the number is even and greater than 2, start checking from the next odd number
	if nb%2 == 0 {
		nb++
	}
	// Loop until a prime number is found
	for !isPrime(nb) {
		nb += 2 // Skip even numbers
	}
	return nb
}

// isPrime checks if a number is prime.
func isPrime(nb int) bool {
	if nb <= 1 {
		return false // 0 and 1 are not prime numbers
	}
	if nb == 2 {
		return true // 2 is prime
	}
	// Check divisibility from 3 up to the square root of nb (without using math package)
	limit := intSqrt(nb)
	for i := 3; i <= limit; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

// intSqrt calculates the integer square root of a number.
func intSqrt(n int) int {
	if n <= 1 {
		return n
	}

	// Perform a binary search for the integer square root
	left, right := 1, n
	var mid int
	for left <= right {
		mid = (left + right) / 2
		square := mid * mid
		if square == n {
			return mid
		} else if square < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
