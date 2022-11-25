/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */
package main

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[low] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[mid] == nums[high] {
				high--
			}
			if nums[mid] == nums[low] {
				low++
			}
		}
	}
	return -1
}

// @lc code=end
