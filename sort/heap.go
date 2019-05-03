package sort

import "github.com/eno314/GoSamples/datastruct"

type HeapSort struct{}

// NewHeapSort is create heap sort struct
func NewHeapSort() Sortable {
	var sortable Sortable = &HeapSort{}
	return sortable
}

// Sort is sort by heap sort
func (h *HeapSort) Sort(target []int) []int {
	heap := datastruct.Heap{Nodes: []int{}}

	for _, value := range target {
		heap.Add(value)
	}

	sorted := []int{}
	for true {
		value := heap.Put()
		if value < 0 {
			break
		}
		sorted = append(sorted, value)
	}

	return sorted
}
