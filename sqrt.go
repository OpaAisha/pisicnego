package piscine

func Sqrt(nb int) int {
	var del int
	if nb <= 0 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		if nb/i == i && nb%i == 0 {
			del = i
		}
	}
	return del
}
