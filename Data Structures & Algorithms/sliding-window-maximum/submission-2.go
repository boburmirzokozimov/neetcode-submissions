func maxSlidingWindow(nums []int, k int) []int {
	out := []int{}
	h := MaxHeap{}
	heap.Init(&h)

	for i := 0; i < len(nums); i++ {
		heap.Push(&h, [2]int{nums[i], i})

		if i >= k - 1 {
			for i - h[0][1] >= k {
				heap.Pop(&h)
			}
			out = append(out, h[0][0])
		}
	}

	return out
}

type MaxHeap [][2]int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i][0] > h[j][0]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}
func (h *MaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
