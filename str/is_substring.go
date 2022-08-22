package str

// KMP match: https://www.youtube.com/watch?v=GTJr8OvyEVQ
func IsSubstring(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	// when needle is blank string return 0
	if m == 0 {
		return 0
	}

	// prefix array, when not match with i in needle, which index to continue
	prefixIndexes := make([]int, m)
	// when the first char is not match, should continue from first char, so i start from 1
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = prefixIndexes[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		prefixIndexes[i] = j
	}

	// iteralte the haystack once
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = prefixIndexes[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	
	return -1
}
