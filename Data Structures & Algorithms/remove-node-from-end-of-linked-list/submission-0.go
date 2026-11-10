/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	curr := dummy
//head = [5], n = 1
//head = [1,2,3,4], n = 2
	fast := dummy 
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	slow := curr
	
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
