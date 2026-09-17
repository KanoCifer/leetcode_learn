package leetcodelearn

func detectCycle(head *ListNode) *ListNode {
	var slow = head
	var fast = head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}

			return slow
		}
	}

	return nil
}
