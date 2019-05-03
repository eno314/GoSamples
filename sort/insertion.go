package sort

type InsertionSort struct{}

// NewInsertionSort is create insertion sort struct
func NewInsertionSort() Sortable {
	var sortable Sortable = &InsertionSort{}
	return sortable
}

// Sort is sort by insertion sort
func (i *InsertionSort) Sort(target []int) []int {
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
