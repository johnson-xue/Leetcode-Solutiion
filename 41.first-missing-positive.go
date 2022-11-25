/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 */
package main

// @lc code=start
func firstMissingPositive(nums []int) int {
	// 数组长度为n，答案可能的取值为[1, n+1]
	// 原地哈希：index-i => number-i+1
	n := len(nums)
	for i := 0; i < n; i++ {
		for 1 <= nums[i] && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return n + 1
}

// @lc code=end
