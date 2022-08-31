package o

func BinarySearch(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length
	for left < right {
		middle := (left + right) / 2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return -1
}
