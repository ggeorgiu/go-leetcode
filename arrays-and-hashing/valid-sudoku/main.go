package main

import (
	"fmt"
)

func main() {
	input := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(input))
}

func isValidSudoku(board [][]byte) bool {
	var rows, cols, squares [9][9]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c == '.' {
				continue
			}

			nr := int(c - '1')
			sq := getSquare(i, j)

			rows[nr][i]++
			cols[nr][j]++
			squares[nr][sq]++

			if rows[nr][i] > 1 || cols[nr][j] > 1 || squares[nr][sq] > 1 {
				return false
			}
		}
	}

	return true
}

func getSquare(x, y int) int {
	nChIdx := x*3 + y/3

	row := nChIdx / 9
	col := nChIdx % 3

	return row*3 + col
}
