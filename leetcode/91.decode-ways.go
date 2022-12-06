/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */
package main

// @lc code=start
//
//	func numDecodings(s string) int {
//		n := len(s)
//		dp := make([]int, n+1)
//		dp[0] = 1
//		for i := 1; i <= n; i++ {
//			if s[i-1] != '0' {
//				dp[i] += dp[i-1]
//			}
//			if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
//				dp[i] += dp[i-2]
//			}
//		}
//		return dp[n]
//	}
func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		if s[i] != '0' {
			dp[i] += dp[i+1]
			if i+1 < n && (s[i]-'0')*10+(s[i+1]-'0') <= 26 {
				dp[i] += dp[i+2]
			}
		}
	}
	return dp[0]
}

// @lc code=end
