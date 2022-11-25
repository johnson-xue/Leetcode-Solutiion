/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */
package main

// @lc code=start
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	maxJump := 0
	for i, v := range nums {
		if i > maxJump {
			return false
		}
		maxJump = max1(maxJump, i+v)
		// if maxJump >= len(nums) {
		// 	return true
		// }
	}
	return true
}
func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
