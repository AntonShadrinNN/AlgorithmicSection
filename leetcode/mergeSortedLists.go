package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result = &ListNode{}
	head := result
	p1, p2 := list1, list2
	for p1.Next != nil && p2.Next != nil {
		if p1.Val < p2.Val {
			result.Val = p1.Val
			result.Next = &ListNode{}
			result = result.Next
			p1 = list1.Next
		} else {
			result.Val = p2.Val
			result.Next = &ListNode{}
			result = result.Next
			p2 = list2.Next
		}
	}
	return head
}
