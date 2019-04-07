package sort

import (
	"github.com/eno314/GoSamples/datastruct"
)

// Bubble is sort by bubble sort.
func Bubble(target []int) []int {
	sorted := append([]int{}, target...)

	lenght := len(sorted)
	for i := 0; i < lenght; i++ {
		for j := lenght; j >= i+2; j-- {
			if sorted[j-1] < sorted[j-2] {
				tmp := sorted[j-1]
				sorted[j-1] = sorted[j-2]
				sorted[j-2] = tmp
			}
		}
	}

	return sorted
}

// Selection is sort by selection sort.
func Selection(target []int) []int {
	sorted := append([]int{}, target...)

	length := len(sorted)
	for i := 0; i < length-1; i++ {
		minValue := sorted[i]
		minIndex := i
		for j := i + 1; j < length; j++ {
			if sorted[j] < minValue {
				minValue = sorted[j]
				minIndex = j
			}
		}
		sorted[minIndex] = sorted[i]
		sorted[i] = minValue
	}
	return sorted
}

// Insertion is sort by insertion sort.
func Insertion(target []int) []int {
	sorted := append([]int{}, target...)

	length := len(sorted)
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if sorted[i] < sorted[j] {
				tmp := sorted[i]
				sorted[i] = sorted[j]
				sorted[j] = tmp
				continue
			}
		}
	}

	return sorted
}

// Heap is sort by heap sort.
func Heap(target []int) []int {
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
