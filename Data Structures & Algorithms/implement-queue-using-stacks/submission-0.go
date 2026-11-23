type MyQueue struct {
	q []int
}

func Constructor() MyQueue {
	return MyQueue{q: []int{}}
}

func (this *MyQueue) Push(x int) {
	this.q = append(this.q, x)
}

func (this *MyQueue) Pop() int {
	front := this.q[0]
	this.q = this.q[1:]
	return front
}

func (this *MyQueue) Peek() int {
	return this.q[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.q) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
