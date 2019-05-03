package sort

import (
	"time"
)

type Sortable interface {
	Sort([]int) []int
}

func SortWithDuration(sortable Sortable, target []int) (sorted []int, duration time.Duration) {
	startTime := time.Now()

	sorted = sortable.Sort(target)
	duration = time.Now().Sub(startTime)
	return
}

func Sortables() []Sortable {
	return []Sortable{
		NewBubbleSort(),
		NewSelectionSort(),
		NewInsertionSort(),
		NewHeapSort(),
		NewMergeSort(),
		NewQuickSort(),
	}
}
