/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

    firstHead := head

	slow, fast := head, head.Next

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tmp := slow.Next
	slow.Next = nil
	curr := tmp
	var prev *ListNode

	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	
	secondHead := prev
	secCurr := secondHead
	firstCurr := firstHead
	mv := firstHead

	for secCurr != nil && firstCurr != nil {
		tmp1 := firstCurr.Next
		tmp2 := secCurr.Next

		mv.Next = firstCurr
		mv = mv.Next
		mv.Next = secCurr
		mv = mv.Next

		firstCurr = tmp1
		secCurr = tmp2
	}

	if firstCurr != nil {
		mv.Next = firstCurr
	} else {
		mv.Next = secCurr
	}

}
