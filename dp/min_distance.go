package dp

import "fmt"

func MinDistance(s1, s2 string) int {
	m, n := len(s1), len(s2)
	// dp[i][j] 存储 s1[0:i] 变成 s2[0:j] 的最小编辑距离
	dp := make([][]int, m+1)

	// 当 j=0 时，最小编辑距离为 i，即删除 i 次
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	// 当 i=0 时，最小编辑距离为 j，即插入 j 次
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j-1], //替换
					dp[i-1][j],   //删除
					dp[i][j-1],   //插入
				) + 1
			}
		}
	}

	fmt.Println(dp)

	return dp[m][n]
}
