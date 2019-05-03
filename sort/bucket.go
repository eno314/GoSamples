package sort

type BucketSort struct{}

// NewBucketSort is create bucket sort struct
func NewBucketSort() Sortable {
	var sortable Sortable = &BucketSort{}
	return sortable
}

// Sort is sort by bucket sort
func (b *BucketSort) Sort(target []int) []int {
	sorted := append([]int{}, target...)

	bucket := make([]int, maxValue(target)+1)
	for _, value := range sorted {
		bucket[value] = bucket[value] + 1
	}

	index := 0
	for value, count := range bucket {
		for i := 0; i < count; i++ {
			sorted[index] = value
			index++
		}
	}

	return sorted
}

func maxValue(target []int) int {
	max := 0
	for _, value := range target {
		if max < value {
			max = value
		}
	}
	return max
}
