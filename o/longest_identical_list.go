package o

func LongestIdenticalList(s string) int {
	res := 0
	dict := make(map[byte]bool, 0)
	right := 0
	length := len(s)
	for left := 0; left < length; {
		for right < length {
			// if right pointer char is exsit
			if _, ok := dict[s[right]]; ok {
				break
			}
			dict[s[right]] = true
			right++
		}

		res = max(res, right-left)

		// move left pointer 1 step
		left++
		// delete the char before cur window
		delete(dict, s[left-1])
	}

	return res
}
