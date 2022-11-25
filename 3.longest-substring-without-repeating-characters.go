/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
package main

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	low, right, result := 0, -1, 0
	var freq [127]int
	for low < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[low]]--
			low++
		}
		result = Max(result, (right + 1 - low))
	}
	return result
}

// @lc code=end
