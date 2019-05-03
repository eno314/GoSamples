package sort

type SelectionSort struct{}

// NewSelectionSort is create selection sort struct
func NewSelectionSort() Sortable {
	var selectionSort Sortable = &SelectionSort{}
	return selectionSort
}

// Sort is sort by selection sort.
func (selection *SelectionSort) Sort(target []int) []int {
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
