func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	h := &MinHeap{}
	heap.Init(h)
	for key, val := range freq {
		p := Pair{
			num: key,
			freq: val,
		}
		heap.Push(h, p)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, 0, k)
	for h.Len() > 0 {
		result = append(result, h.Pop().(Pair).num)
	}

	return result
}

type Pair struct {
	num  int
	freq int
}

type MinHeap []Pair

func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}
func (h *MinHeap) Pop() any {
	old := *h
	n := old[len(old)-1]
	*h = old[:len(old)-1]
	return n
}

