func maxSlidingWindow(nums []int, k int) []int {
    maxHeap := &MaxHeap{}

	heap.Init(maxHeap)

	out := []int{}

	for i := 0; i < len(nums); i++ {
		heap.Push(maxHeap, [2]int{nums[i], i})
		if i >= k - 1 {
			for (*maxHeap)[0][1] <= i - k {
				heap.Pop(maxHeap)
			}
			out = append(out, (*maxHeap)[0][0])
		}
	}

	return out
}

type MaxHeap [][2]int

func (h MaxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Len() int { return len(h) }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}


