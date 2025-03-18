package piscine

import "fmt"

func allPositive(x, y int) bool {
	return x > 0 && y > 0
}

func QuadA(x, y int) {
	if !allPositive(x, y) {
		return
	}

	for j := 1; j <= y; j++ {
		for i := 1; i <= x; i++ {
			switch {
			case (i == 1 || i == x) && (j == 1 || j == y):
				fmt.Print("o")
			case j == 1 || j == y:
				fmt.Print("-")
			case i == 1 || i == x:
				fmt.Print("|")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func QuadB(x, y int) {
	if !allPositive(x, y) {
		return
	}

	for j := 1; j <= y; j++ {
		for i := 1; i <= x; i++ {
			switch {
			case (i == 1 || i == x) && (j == 1 || j == y):
				if (i == 1 && j == 1) || (i == x && j == y && x != 1 && y != 1) {
					fmt.Print("/")
				} else {
					fmt.Print("\\")
				}
			case j == 1 || j == y || i == 1 || i == x:
				fmt.Print("*")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func QuadC(x, y int) {
	if !allPositive(x, y) {
		return
	}

	for j := 1; j <= y; j++ {
		for i := 1; i <= x; i++ {
			switch {
			case (i == 1 || i == x) && (j == 1 || j == y):
				if j == 1 {
					fmt.Print("A")
				} else {
					fmt.Print("C")
				}
			case (j == 1 || j == y) || (i == 1 || i == x):
				fmt.Print("B")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func QuadD(x, y int) {
	if !allPositive(x, y) {
		return
	}

	for j := 1; j <= y; j++ {
		for i := 1; i <= x; i++ {
			switch {
			case (i == 1 || i == x) && (j == 1 || j == y):
				if i == 1 {
					fmt.Print("A")
				} else {
					fmt.Print("C")
				}
			case (j == 1 || j == y) || (i == 1 || i == x):
				fmt.Print("B")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func QuadE(x, y int) {
	if !allPositive(x, y) {
		return
	}

	for j := 1; j <= y; j++ {
		for i := 1; i <= x; i++ {
			switch {
			case (i == 1 || i == x) && (j == 1 || j == y):
				if (i == 1 && j == 1) || (i == x && j == y && x != 1 && y != 1) {
					fmt.Print("A")
				} else {
					fmt.Print("C")
				}
			case j == 1 || j == y || i == 1 || i == x:
				fmt.Print("B")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
