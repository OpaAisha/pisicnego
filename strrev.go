package piscine

func StrRev(s string) string {
	runes := []rune(s) // Convert string to rune slice
	start := 0
	end := len(runes) - 1

	// Reverse the slice of runes
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	return string(runes)
}
