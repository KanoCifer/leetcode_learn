package leetcodelearn

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

// 复制链表
func reverseList(head *ListNode) *ListNode {
    var cur *ListNode
    for head != nil {
        cur = &ListNode{
            Val:  head.Val,
            Next: cur,
        }
        head = head.Next
    }
    return cur
}

// 原地操作
func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev

		prev = head
		head = next
	}
	return prev
}
