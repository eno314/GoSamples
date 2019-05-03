package sort

import (
	"reflect"
	"testing"
)

type Case struct {
	input  []int
	expect []int
}

func TestSort(t *testing.T) {
	for _, sortable := range Sortables() {
		for _, c := range cases() {
			actual := sortable.Sort(c.input)

			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("[%v]Sort(%v), expect %v, actual %v\n",
					reflect.TypeOf(sortable), c.input, c.expect, actual)
			}
		}
	}
}

func cases() []Case {
	return []Case{
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 3, 2}, []int{1, 2, 3}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		{[]int{2, 3, 1}, []int{1, 2, 3}},
		{[]int{3, 1, 2}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 2, 2}, []int{1, 2, 2}},
		{[]int{2, 1, 2}, []int{1, 2, 2}},
		{[]int{2, 2, 1}, []int{1, 2, 2}},
	}
}
