package sort

import (
	"math/rand"
	"time"

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

// Merge is sort by Merge sort.
func Merge(target []int) []int {
	targetLenght := len(target)

	// 配列の前半・後半に分ける
	firstSize := targetLenght - (targetLenght / 2)
	first := target[firstSize:]
	second := target[:firstSize]
	if targetLenght > 2 {
		// 2件以上の場合は再帰的にマージソートを繰り返す
		first = Merge(first)
		second = Merge(second)
	}

	// 前半・後半の先頭を比較しながら全体をソートする
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

// Quick is sort by Quick sort.
func Quick(target []int) []int {
	if len(target) < 2 {
		return append([]int{}, target...)
	}

	rand.Seed(time.Now().UnixNano())
	pivod := target[rand.Intn(len(target)-1)]

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

	first = Quick(first)
	second = Quick(second)

	return append(append(first, sames...), second...)
}
