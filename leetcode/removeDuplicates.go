package leetcode

func RemoveDuplicates(nums []int) int {
	alreadySorted := 1
	i := 1
	for i < len(nums) {
		if nums[i] != nums[alreadySorted-1] {
			nums[alreadySorted] = nums[i]
			alreadySorted++
			i++
		} else {
			for i < len(nums) && nums[i] == nums[alreadySorted-1] {
				i++
			}
		}
	}

	return alreadySorted
}
