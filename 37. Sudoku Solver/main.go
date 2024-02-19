package main

import (
// "fmt"
)

/*

Determine if a 9 x 9 Sudoku board is valid.
Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.

Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]

*/

// func isValidSudoku(board [][]byte) bool {
// 	// existMap := make(map[int]int)

// 	for j := 0; j < len(board); j++ {
// 		// c := board[i : i+3]
// 		RowExistMap := make(map[byte]int)
// 		ColumnExistMap := make(map[byte]int)
// 		BlockExistMap := make(map[byte]int)

// 		for i := 0; i < len(board[j]); i++ {
// 			// fmt.Println(board[j/3+(i%3)*3][j%3+(i%3)*3])
// 			// fmt.Println("block:", p)
// 			// fmt.Println("x: ", j/3+(i%3)*3, " y: ", j%3+(i%3)*3)

// 			// // BlockExistMap[board[i][j]] &&
// 			if board[j/3+(i%3)*3][j%3+(i%3)*3] != 0 && BlockExistMap[board[j/3+(i%3)*3][j%3+(i%3)*3]] == 0 {

// 				BlockExistMap[board[j/3+(i%3)*3][j%3+(i%3)*3]]++
// 			} else if board[j/3+(i%3)*3][j%3+(i%3)*3] != 0 {
// 				return false
// 			}

// 			if board[i][j] != 0 && RowExistMap[board[i][j]] == 0 {
// 				RowExistMap[board[i][j]]++
// 			} else if board[i][j] != 0 {
// 				return false
// 			}

// 			if board[j][i] != 0 && ColumnExistMap[board[j][i]] == 0 {
// 				ColumnExistMap[board[j][i]]++
// 			} else if board[j][i] != 0 {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
func isValidSudoku(board [][]byte) bool {
	var row, col, block [10][10]bool
	for i := range board {
		for j, c := range board[i] {
			if c == '.' {
				continue
			}
			if col[j][c-'0'] || row[i][c-'0'] || block[i/3+j/3*3][c-'0'] {
				return false
			}
			col[j][c-'0'] = true
			row[i][c-'0'] = true
			block[i/3+j/3*3][c-'0'] = true
		}
	}
	return true
}

// func isValidSudoku(board [][]byte) bool {
// 	var row, col, block [10]int
// 	for i := range board {
// 		for j, c := range board[i] {
// 			if c == '.' {
// 				continue
// 			}
// 			v := 1 << int(c-'0')
// 			if (col[j]&v)|(row[i]&v)|(block[i/3+j/3*3]&v) > 0 {
// 				return false
// 			}
// 			col[j] |= v
// 			row[i] |= v
// 			block[i/3+j/3*3] |= v
// 		}
// 	}
// 	return true
// }
func main() {
	nums := [][]byte{
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

	isValidSudoku(nums)
}

// [['5','3','.','.','7','.','.','.','.']
// ,['6','.','.','1','9','5','.','.','.']
// ,['.','9','8','.','.','.','.','6','.']
// ,['8','.','.','.','6','.','.','.','3']
// ,['4','.','.','8','.','3','.','.','1']
// ,['7','.','.','.','2','.','.','.','6']
// ,['.','6','.','.','.','.','2','8','.']
// ,['.','.','.','4','1','9','.','.','5']
// ,['.','.','.','.','8','.','.','7','9']]
