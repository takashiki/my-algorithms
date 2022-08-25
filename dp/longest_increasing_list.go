package dp

func LongestIncreasingList(nums []int) int {
	// build an array of tail numbers as dp array
	// nums in this array is increasing, so we can search by Dichotomy 
	tails := make([]int, len(nums))
	res := 0

	for _, num := range nums {
		// representing left and right index of search bound in dp array
		left, right := 0, res
		for left < right {
			middle := (left + right) / 2
			// if middle value is less than cur num, search it in right half
			if tails[middle] < num {
				left = middle + 1
			} else {
				// otherwise
				right = middle
			}
		}
		// update tail num when search complete
		tails[left] = num
		// if right equals to res, it mean cur num is max num of longest sublist
		if right == res {
			res++
		}
	}

	return res
}
