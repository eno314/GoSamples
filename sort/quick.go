package sort

type QuickSort struct{}

// NewQuickSort is create quick sort struct
func NewQuickSort() Sortable {
	var sortable Sortable = &QuickSort{}
	return sortable
}

// Sort is sort by quick sort
func (q *QuickSort) Sort(target []int) []int {
	return quickSort(target)
}

func quickSort(target []int) []int {
	if len(target) < 2 {
		return append([]int{}, target...)
	}

	// rand.Seed(time.Now().UnixNano())
	// pivod := target[rand.Intn(len(target)-1)]
	pivod := target[0]

	first := []int{}
	second := []int{}
	sames := []int{}
	for _, v := range target {
		if v < pivod {
			first = append(first, v)
		} else if v > pivod {
			second = append(second, v)
		} else {
			sames = append(sames, v)
		}
	}

	first = quickSort(first)
	second = quickSort(second)

	return append(append(first, sames...), second...)
}
