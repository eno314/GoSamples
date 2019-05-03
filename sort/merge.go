package sort

type MergeSort struct{}

// NewMergeSort is create merge sort struct
func NewMergeSort() Sortable {
	var sortable Sortable = &MergeSort{}
	return sortable
}

// Sort is sort by merge sort
func (m *MergeSort) Sort(target []int) []int {
	// 再帰的に呼び出したいので、内部メソッドに分ける
	return mergeSort(target)
}

func mergeSort(target []int) []int {
	targetLenght := len(target)

	// 配列の前半・後半に分ける
	firstSize := targetLenght - (targetLenght / 2)
	first := target[firstSize:]
	second := target[:firstSize]
	if targetLenght > 2 {
		// 2件以上の場合は再帰的にマージソートを繰り返す
		first = mergeSort(first)
		second = mergeSort(second)
	}

	return merge(first, second)
}

// 2つの配列の先頭を比較しながら全体をソートした1つの配列にして返す
func merge(first []int, second []int) []int {
	sorted := []int{}
	firstIndex := 0
	secondIndex := 0
	for firstIndex < len(first) && secondIndex < len(second) {
		if first[firstIndex] < second[secondIndex] {
			sorted = append(sorted, first[firstIndex])
			firstIndex++
		} else {
			sorted = append(sorted, second[secondIndex])
			secondIndex++
		}
	}
	if firstIndex < len(first) {
		sorted = append(sorted, first[firstIndex:]...)
	}
	if secondIndex < len(second) {
		sorted = append(sorted, second[secondIndex:]...)
	}

	return sorted
}
