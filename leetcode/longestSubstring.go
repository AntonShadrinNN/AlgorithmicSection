package leetcode

import "strings"

func LengthOfLongestSubstring(s string) int {
	maxLength := 0
	subString := ""
	for _, val := range s {
		if !strings.Contains(subString, string(val)) {
			subString += string(val)
			if len(subString) > maxLength {
				maxLength = len(subString)
			}
		} else {
			subString = subString[strings.LastIndex(subString, string(val))+1:] + string(val)
		}
	}
	return maxLength
}
