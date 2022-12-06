/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */
package main

// @lc code=start
func largestRectangleArea(heights []int) int {
	maxArea := 0
	n := len(heights) + 2
	getHeight := func(i int) int {
		if i == 0 || n == i+1 {
			return 0
		}
		return heights[i-1]
	}
	st := make([]int, 0, n/2)
	for i := 0; i < n; i++ {
		for len(st) > 0 && getHeight(st[len(st)-1]) > getHeight(i) {
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			maxArea = max3(maxArea, getHeight(idx)*(i-st[len(st)-1]-1))
		}
		st = append(st, i)
	}
	return maxArea
}
func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
