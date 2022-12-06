/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */
package main

import "sort"

// @lc code=start
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res, used, path := [][]int{}, make([]bool, len(nums)), []int{}
	sort.Ints(nums)
	generatePermutationUnique(nums, 0, path, &used, &res)
	return res
}
func generatePermutationUnique(nums []int, index int, path []int, used *[]bool, res *[][]int) {
	if index == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			if i > 0 && !(*used)[i-1] && nums[i] == nums[i-1] {
				continue
			}
			(*used)[i] = true
			path = append(path, nums[i])
			generatePermutationUnique(nums, index+1, path, used, res)
			path = path[:len(path)-1]
			(*used)[i] = false
		}
	}
	return
}

// @lc code=end
