/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, l int, r int) *ListNode {
	dummy := &ListNode{Next: head}
	curr := dummy
	var prev *ListNode

	for i := 0; i < l; i++ {
		prev = curr
		curr = curr.Next
	}
	stopped := prev
	firstHead := curr

	for l < r {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
		l++
	}
	tmp := curr.Next
	curr.Next = prev
	firstHead.Next = tmp
	stopped.Next = curr

	return dummy.Next
}
