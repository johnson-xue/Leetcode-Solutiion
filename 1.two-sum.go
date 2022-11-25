/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package main

// @lc code=start
//package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

// @lc code=end
