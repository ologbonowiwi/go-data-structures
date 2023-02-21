package sort

func QuickSort(array []int) {
	// Check if the array is empty or has only one element
	if len(array) <= 1 {
		// An array of length 0 or 1 is already sorted
		return
	}

	// Choose the pivot as the first element of the array
	pivot := array[0]

	// Create two empty slices to hold elements less than and greater than the pivot
	var left []int
	var right []int

	// Start iterating from the second element in the array
	for _, value := range array[1:] {
		if value < pivot {
			// If the value is less than the pivot, add it to the left slice
			left = append(left, value)
		} else {
			// If the value is greater than or equal to the pivot, add it to the right slice
			right = append(right, value)
		}
	}

	// Sort the left and right slices recursively
	QuickSort(left)
	QuickSort(right)

	// Merge the sorted left and right slices with the pivot to form the final sorted array
	leftWithPivot := append(left, pivot)
	copy(array, append(leftWithPivot, right...))
}
