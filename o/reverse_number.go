package o

func ReverseNumber(num int) int {
	res := 0
	for num > 0 {
		n := num % 10
		res = res*10 + n
		num = num / 10
	}

	return res
}
