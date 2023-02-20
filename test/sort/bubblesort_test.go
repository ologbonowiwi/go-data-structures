package sort_test

import (
	"reflect"
	"testing"

	"github.com/ologbonowiwi/go-data-structures/pkg/sort"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
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

	for _, test := range tests {
		sort.BubbleSort(test.input)

		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("BubbleSort failed, expected %v but got %v", test.expected, test.input)
		}
	}
}
