/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */
package main

// @lc code=start
func searchRange(nums []int, target int) []int {
	start, end, mid, res := 0, len(nums), 0, []int{}
	if end == 0 {
		return []int{-1, -1}
	}
	for start < end {
		mid = start + (end-start)>>1
		if nums[mid] == target {
			for i := mid; i >= start; i-- {
				if nums[i] != target {
					res = append(res, i+1)
					break
				}
				if i == start {
					res = append(res, start)
				}
			}
			realEnd := 0
			if end == len(nums) {
				realEnd = end - 1
			} else {
				realEnd = end
			}
			for j := mid; j <= realEnd; j++ {
				if nums[j] != target {
					res = append(res, j-1)
					break
				}
				if j == realEnd {
					res = append(res, realEnd)
				}
			}
			return res
		} else if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		}
	}
	if start == end {
		if start == len(nums) {
			start--
		}
		if nums[start] == target {
			return []int{start, start}
		} else {
			return []int{-1, -1}
		}
	}
	return []int{-1, -1}
}

// @lc code=end
