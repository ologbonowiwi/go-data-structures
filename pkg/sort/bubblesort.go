package sort

func BubbleSort(array []int) {
	iterable := len(array) - 1

	for iterable != 0 {
		for i := 0; i < iterable; i++ {
			a, b := i, i+1
			if array[a] > array[b] {
				array[a], array[b] = array[b], array[a]
			}
		}

		iterable = iterable - 1
	}
}
