/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */
package main

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 1 {
		return
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			swap(matrix, i, j)
		}
		matrix[i] = reverse1(matrix[i])
	}
}

// swap changes original slice's i,j position
func swap(nums [][]int, i, j int) {
	nums[i][j], nums[j][i] = nums[j][i], nums[i][j]
}

// reverses a row of image, matrix[i]
func reverse1(nums []int) []int {
	var lp, rp = 0, len(nums) - 1
	for lp < rp {
		nums[lp], nums[rp] = nums[rp], nums[lp]
		lp++
		rp--
	}
	return nums
}

// @lc code=end
