package o

// (220, 284) is the smallest AP
// 220 -> 1, 2, 4, 5... 110, 220 -> 284
// 284 -> 1, 2, 4, ... 142, 284 -> 220
func AmicablePair(n int) [][]int {
	res := make([][]int, 0)
	pairMap := make(map[int]int, 0)
	for i := 1; i <= n; i++ {
		amicableNum := calcAmicableNum(i)
		// amicable pair num should not equals to it self
		if amicableNum == n {
			continue
		}

		// if amicableNum > n, i has no amicable num in n range
		if amicableNum > n {
			continue
		}

		// if the amicable num of current amicable num equals to itself
		if value, exsit := pairMap[amicableNum]; exsit && value == i {
			r := []int{i, amicableNum}
			res = append(res, r)
			continue
		}

		pairMap[i] = amicableNum
	}

	return res
}

func calcAmicableNum(n int) int {
	sum := 1
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	return sum
}
