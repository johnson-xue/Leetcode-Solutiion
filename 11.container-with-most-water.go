/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
package main

// @lc code=start
func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		temp := high * width
		if temp > max {
			max = temp
		}
	}
	return max
}

// @lc code=end
