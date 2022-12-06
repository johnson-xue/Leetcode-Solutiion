/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */
package main

// @lc code=start
func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		mid := (l + r + 1) / 2
		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}

// @lc code=end
