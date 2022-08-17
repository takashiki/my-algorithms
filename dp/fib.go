package dp

func Fib(n int) int {
	// 初始化当 n=1 时，返回 1
	fibN := 1
	// 这两个变量用来存储 dp[i-2] 和 dp[i-1]
	fibNpp, fibNp := 0, 1
	// 如果 n=1 则不会进入循环，直接返回1，n>1时就是一个常规的dp问题
	for i := 2; i <= n; i++ {
		fibN = fibNpp + fibNp
		fibNpp = fibNp
		fibNp = fibN
	}

	return fibN
}
