/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */
package main

import "math"

// @lc code=start
func divide(dividend int, divisor int) int {
	res, sign := 0, -1
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if dividend == math.MaxInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		sign = 1
	}
	if dividend > math.MaxInt32 {
		dividend = math.MaxInt32
	}
	res = binarySearchQuotient(0, abs(dividend), abs(divisor), abs(dividend))
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	if res < math.MinInt32 {
		return math.MinInt32
	}
	return sign * res
}
func binarySearchQuotient(low, high, val, dividend int) int {
	quotient := low + (high-low)>>1
	if ((quotient+1)*val >= dividend && quotient*val < dividend) || ((quotient+1)*val > dividend && quotient*val <= dividend) {
		if (quotient+1)*val == dividend {
			return quotient + 1
		}
		return quotient
	}
	if (quotient+1)*val > dividend && quotient*val > dividend {
		return binarySearchQuotient(low, quotient-1, val, dividend)
	}
	if (quotient+1)*val < dividend && quotient*val < dividend {
		return binarySearchQuotient(quotient+1, high, val, dividend)
	}
	return 0
}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// @lc code=end
