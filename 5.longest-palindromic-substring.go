/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
package main

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, 0, 0, 0
	for left < len(s) {
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}

// @lc code=end
