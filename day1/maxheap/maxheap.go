package maxheap

// A binary heap of ints
// Implements the container/heap interface and is a light modification of
// the IntHeap example present in the documentation:
// https://pkg.go.dev/container/heap#pkg-functions
type MaxHeap []int

// Len returns the number of elements in the heap
func (h MaxHeap) Len() int { return len(h) }

// Less returns whether the element at index j is less than the element at
// index i. Comparison for sorting in descending order.
func (h MaxHeap) Less(i, j int) bool { return h[j] < h[i] }

// Swap swaps the elements at position i and j.
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push pushes the element x onto the heap.
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the minimum element from the heap.
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
