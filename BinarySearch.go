package main

func binarySearchMatrix(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 二分查找第一个与target相同的元素，时间复杂度 O(logn)
func searchFirstEqualRlement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// 二分查找最后一个与target相同的元素，时间复杂度 O(logn)
func searchLastEqualRlement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

// 第一个大于等于
func binarySearchFistGreaterElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 最后一个小于等于
func searchLastLessElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] <= target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
