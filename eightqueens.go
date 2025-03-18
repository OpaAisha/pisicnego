package piscine

import (
	"github.com/01-edu/z01"
)

func isSafe(board []int, row, col int) bool {
	for i := 0; i < col; i++ {
		if board[i] == row || board[i]-i == row-col || board[i]+i == row+col {
			return false
		}
	}
	return true
}

func solve(board []int, col int) {
	if col == 8 {
		// Печатаем результат в правильном формате
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(board[i] + '1')) // Преобразуем в символ от '1' до '8'
		}
		z01.PrintRune('\n')
		return
	}

	for row := 0; row < 8; row++ {
		if isSafe(board, row, col) {
			board[col] = row
			solve(board, col+1)
		}
	}
}

func EightQueens() {
	// Используем обычный массив, так как это может требоваться платформой
	var board [8]int
	// Передаем массив как срез в функцию solve
	solve(board[:], 0)
}
