package sort

import (
	"fmt"
	"reflect"
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

// ShowCalcDurationBySort is show to calcuration duration by sort
func ShowCalcDurationBySort(sortable Sortable, target []int) {
	_, duration := SortWithDuration(sortable, target)
	fmt.Printf("call took %v to run by %v\n", duration, reflect.TypeOf(sortable))
}

// Sortables is return types implements Sortable
func Sortables() []Sortable {
	return []Sortable{
		NewBubbleSort(),
		NewSelectionSort(),
		NewInsertionSort(),
		NewHeapSort(),
		NewMergeSort(),
		NewQuickSort(),
		NewBucketSort(),
	}
}
