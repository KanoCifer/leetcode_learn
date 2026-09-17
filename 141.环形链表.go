package leetcodelearn

func hasCycle(head *ListNode) bool {
	var slow = head
	var fast = head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
