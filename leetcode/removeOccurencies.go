package leetcode

func RemoveElement(nums []int, val int) int {
	unique := len(nums)
	i := 0
	for i != unique {
		if nums[i] == val {
			nums[unique-1], nums[i] = nums[i], nums[unique-1]
			unique--
		} else {
			i++
		}
	}
	return unique
}
