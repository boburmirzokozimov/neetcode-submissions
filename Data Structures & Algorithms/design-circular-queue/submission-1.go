type Node struct {
	Val  int
	Next *Node
}

type MyCircularQueue struct {
	Head, Tail *Node
	Volume     int
	Filled     int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{Volume: k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	node := &Node{Val: value}
	if this.Tail == nil { // пустая: узел и голова, и хвост
		this.Head = node
	} else {
		this.Tail.Next = node
	}
	this.Tail = node
	this.Filled++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.Head = this.Head.Next
	if this.Head == nil { // удалили последний элемент
		this.Tail = nil
	}
	this.Filled--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Head.Val
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Tail.Val
}

func (this *MyCircularQueue) IsEmpty() bool { return this.Filled == 0 }
func (this *MyCircularQueue) IsFull() bool  { return this.Filled == this.Volume }