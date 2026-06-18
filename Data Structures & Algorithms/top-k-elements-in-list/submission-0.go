type Item struct {
	Value int
	Count int
}

type MinHeap []Item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i int, j int) bool {
	return h[i].Count < h[j].Count
} 

func (h MinHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	i, ok := x.(Item)
	if !ok {
		panic("Error")
	}

	*h = append(*h, i)
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	
	i := old[len(old)-1]
	*h = old[:len(old)-1]
	
	return i
}

func topKFrequent(nums []int, k int) []int {
	mem := make(map[int]int)
	for _, v := range nums {
		mem[v]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for num, count := range mem {
		heap.Push(h, Item{num, count})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, 0, k)
	for _, v := range *h {
		res = append(res, v.Value)
	}

	return res
}
