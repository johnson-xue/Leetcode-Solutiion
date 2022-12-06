/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */
package main

// @lc code=start
// 解法1
// const (
// 	right = iota
// 	down
// 	left
// 	up
// )

// var dirs = []int{right, down, left, up}

// func spiralOrder(matrix [][]int) []int {
// 		m := len(matrix)
// if m == 0 {
// 	return nil
// }

// n := len(matrix[0])
//
//	if n == 0 {
//		return nil
//	}
//
//		res := make([]int, 0, m*n)
//		l, r, t, b := 0, n-1, 0, m-1
//		for cnt := 0; l <= r && t <= b; cnt++ {
//			switch dirs[cnt%len(dirs)] {
//			case right:
//				for i, j := t, l; j <= r; j++ {
//					res = append(res, matrix[i][j])
//				}
//				t++
//			case down:
//				for i, j := t, r; i <= b; i++ {
//					res = append(res, matrix[i][j])
//				}
//				r--
//			case left:
//				for i, j := b, r; j >= l; j-- {
//					res = append(res, matrix[i][j])
//				}
//				b--
//			case up:
//				for i, j := b, l; i >= t; i-- {
//					res = append(res, matrix[i][j])
//				}
//				l++
//			}
//		}
//		return res
//	}
//
// 解法二
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	left, right, top, down := 0, n-1, 0, m-1
	count, sum := 0, m*n
	res := []int{}
	for count < sum {
		i, j := top, left
		for j <= right && count < sum {
			res = append(res, matrix[i][j])
			count++
			j++
		}
		i, j = top+1, right
		for i <= down && count < sum {
			res = append(res, matrix[i][j])
			count++
			i++
		}
		i, j = down, right-1
		for j >= left && count < sum {
			res = append(res, matrix[i][j])
			count++
			j--
		}
		i, j = down-1, left
		for i > top && count < sum {
			res = append(res, matrix[i][j])
			count++
			i--
		}
		left, right, top, down = left+1, right-1, top+1, down-1
	}
	return res
}

// @lc code=end
