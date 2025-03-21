package piscine

func SortIntegerTable(table []int) {
	n := len(table)

	// Implementing Bubble Sort
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if table[j] > table[j+1] {
				// Swap the elements
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
