package sort_test

import (
	"reflect"
	"testing"

	"github.com/ologbonowiwi/go-data-structures/pkg/sort"
)

var TEST_CASES = []struct {
	input    []int
	expected []int
}{
	{[]int{64, 34, 25, 12, 22, 11, 90}, []int{11, 12, 22, 25, 34, 64, 90}},
	{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{[]int{7, 6, 8, 9, 10}, []int{6, 7, 8, 9, 10}},
	{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},
	{[]int{3, 3, 3, 3, 3}, []int{3, 3, 3, 3, 3}},
	{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
}

func checkSorting(t *testing.T, sortFunc func([]int), name string) {
	for _, test := range TEST_CASES {
		sortFunc(test.input)

		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("%s failed, expected %v but got %v", name, test.expected, test.input)
		}
	}
}

func TestBubbleSort(t *testing.T) {
	checkSorting(t, sort.BubbleSort, "BubbleSort")
}

func TestQuickSort(t *testing.T) {
	checkSorting(t, sort.QuickSort, "QuickSort")
}
