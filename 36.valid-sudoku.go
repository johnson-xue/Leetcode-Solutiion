/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */
package main

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	rowbuf, colbuf, boxbuf := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowbuf[i] = make([]bool, 9)
		colbuf[i] = make([]bool, 9)
		boxbuf[i] = make([]bool, 9)
	}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				nums := board[r][c] - '0' - byte(1)
				if rowbuf[r][nums] || colbuf[c][nums] || boxbuf[r/3*3+c/3][nums] {
					return false
				}
				rowbuf[r][nums] = true
				colbuf[c][nums] = true
				boxbuf[r/3*3+c/3][nums] = true
			}
		}
	}
	return true
}

// 判断行 row
// for i := 0; i < 9; i++ {
// 	tmp := [10]int{}
// 	for j := 0; j < 9; j++ {
// 		cellVal := board[i][j : j+1]
// 		if string(cellVal) != "." {
// 			index, _ := strconv.Atoi(string(cellVal))
// 			if index > 9 || index < 1 {
// 				return false
// 			}
// 			if tmp[index] == 1 {
// 				return false
// 			}
// 			tmp[index] = 1
// 		}
// 	}
// }
// // 判断列 column
// for i := 0; i < 9; i++ {
// 	tmp := [10]int{}
// 	for j := 0; j < 9; j++ {
// 		cellVal := board[j][i]
// 		if string(cellVal) != "." {
// 			index, _ := strconv.Atoi(string(cellVal))
// 			if index > 9 || index < 1 {
// 				return false
// 			}
// 			if tmp[index] == 1 {
// 				return false
// 			}
// 			tmp[index] = 1
// 		}
// 	}
// }
// // 判断 9宫格 3X3 cell
// for i := 0; i < 3; i++ {
// 	for j := 0; j < 3; j++ {
// 		tmp := [10]int{}
// 		for ii := i * 3; ii < i*3+3; ii++ {
// 			for jj := j * 3; jj < j*3+3; jj++ {
// 				cellVal := board[ii][jj]
// 				if string(cellVal) != "." {
// 					index, _ := strconv.Atoi(string(cellVal))
// 					if tmp[index] == 1 {
// 						return false
// 					}
// 					tmp[index] = 1
// 				}
// 			}
// 		}
// 	}
// }
// return true
// }

// @lc code=end
