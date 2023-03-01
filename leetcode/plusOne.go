package leetcode

func PlusOne(digits []int) []int {
	overflow := 1
	result := make([]int, len(digits)+1)
	for i := len(digits) - 1; i >= 0; i-- {
		result[i+1] = (digits[i] + overflow) % 10
		if digits[i]+overflow == 10 {
			overflow = 1
		} else {
			overflow = 0
		}
	}

	if overflow != 0 {
		result[0] = 1
		return result
	}
	return result[1:]
}
