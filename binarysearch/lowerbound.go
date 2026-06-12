package main

func LowerBound(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	ans := len(arr)

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] >= target {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
