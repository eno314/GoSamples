package sort

// BubbleSort is sort by BubbleSort
func BubbleSort(target []int) []int {
	sorted := append([]int{}, target...)

	lenght := len(sorted)
	for i := 0; i < lenght; i++ {
		for k := lenght; k >= i+2; k-- {
			if sorted[k-1] < sorted[k-2] {
				tmp := sorted[k-1]
				sorted[k-1] = sorted[k-2]
				sorted[k-2] = tmp
			}
		}
	}

	return sorted
}
