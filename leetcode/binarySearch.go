package leetcode

func SearchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
