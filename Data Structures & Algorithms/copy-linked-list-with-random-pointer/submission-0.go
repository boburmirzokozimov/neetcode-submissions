/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    oldToCopy := map[*Node]*Node{nil: nil}

	curr := head 
	for curr != nil {
		copy := &Node{Val: curr.Val}
		oldToCopy[curr] = copy
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		copy := oldToCopy[curr]
		copy.Next = oldToCopy[curr.Next]
		copy.Random = oldToCopy[curr.Random]
		curr = curr.Next
	}

	return oldToCopy[head]
}
