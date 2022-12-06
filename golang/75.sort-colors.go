/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */
package main

// @lc code=start
func sortColors(nums []int) {
	zero, one := 0, 0
	for i, n := range nums {
		nums[i] = 2
		if n <= 1 {
			nums[one] = 1
			one++
		}
		if n == 0 {
			nums[zero] = 0
			zero++
		}
	}
}

// @lc code=end
