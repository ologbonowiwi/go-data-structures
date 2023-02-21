package sort

func MergeSort(array []int) {
	// If the array has less than one element, you can end the function here
	if len(array) <= 1 {
		return
	}

	middle := len(array) / 2

	left := array[:middle]
	right := array[middle:]

	// merge sort both arrays, left and right
	MergeSort(left)
	MergeSort(right)

	// index of left array
	l := 0
	// index of right array
	r := 0
	// index of merged array
	m := 0

	// create an empty array with the size of the original array
	merged := make([]int, len(array))

	// will run while I didn't finished to iterate both arrays
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			merged[m] = left[l]
			l++
		} else {
			merged[m] = right[r]
			r++
		}
		m++
	}

	// iterate left array if remains anything
	for l < len(left) {
		merged[m] = left[l]
		l++
		m++
	}

	// iterate right array if remains anything
	for r < len(right) {
		merged[m] = right[r]
		r++
		m++
	}

	// copy merged array to the original array
	copy(array, merged)
}
