package sort

type BubbleSort struct{}

// NewBubbleSort is create bubble sort struct
func NewBubbleSort() Sortable {
	var bubbleSort Sortable = &BubbleSort{}
	return bubbleSort
}

// Sort is sort by bubble sort
func (bubble *BubbleSort) Sort(target []int) []int {
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
