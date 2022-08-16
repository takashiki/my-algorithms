package str

func RevertString(s string) string {
	arr := []rune(s)
	for left, right := 0, len(arr)-1; left < right; left++ {
		arr[left], arr[right] = arr[right], arr[left]
		right--
	}

	return string(arr)
}
