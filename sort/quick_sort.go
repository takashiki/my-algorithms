package sort

func QuickSort(values []int, leftBound, rightBound int) {
	if leftBound >= rightBound {
		return
	}

	pivot := values[leftBound]
	leftSearch := leftBound
	rightSearch := rightBound

	for {
		for values[leftSearch] < pivot {
			leftSearch++
		}
		for values[rightSearch] > pivot {
			rightSearch--
		}
		if leftSearch >= rightSearch {
			break
		}
		values[leftSearch], values[rightSearch] = values[rightSearch], values[leftSearch]
	}

	QuickSort(values, leftBound, leftSearch-1)
	QuickSort(values, rightSearch+1, rightBound)
}
