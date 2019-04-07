package datastruct

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		init []int
		in   int
		want []int
	}{
		{[]int{}, 1, []int{1}},
		{[]int{1}, 2, []int{1, 2}},
		{[]int{2}, 1, []int{1, 2}},
		{[]int{1, 2}, 3, []int{1, 2, 3}},
		{[]int{1, 3}, 2, []int{1, 3, 2}},
		{[]int{2, 3}, 1, []int{1, 3, 2}},
	}

	for _, c := range cases {
		// initialize
		heap := Heap{c.init}

		// add in to heap
		heap.Add(c.in)

		if !reflect.DeepEqual(heap.Nodes, c.want) {
			t.Errorf("heap.Add(%v) == %v, want %v", c.in, heap.Nodes, c.want)
		}
	}
}

func TestPut(t *testing.T) {
	cases := []struct {
		init []int
		want int
		last []int
	}{
		{[]int{1, 2, 3, 4}, 1, []int{2, 4, 3}},
		{[]int{1, 3, 2, 4}, 1, []int{2, 3, 4}},
		{[]int{1, 2, 3}, 1, []int{2, 3}},
		{[]int{1, 3, 2}, 1, []int{2, 3}},
		{[]int{1, 2}, 1, []int{2}},
		{[]int{1}, 1, []int{}},
		{[]int{}, -1, []int{}},
	}

	for _, c := range cases {
		// initialize
		heap := Heap{c.init}

		// add in to heap
		got := heap.Put()

		if got != c.want || !reflect.DeepEqual(heap.Nodes, c.last) {
			t.Errorf("heap.Put == %v, want %v, last == %v, lastWant %v", got, c.want, heap.Nodes, c.last)
		}
	}
}
