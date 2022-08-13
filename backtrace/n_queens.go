package backtrace

var nQueensResults [][][]bool

func NQueens(n int) [][][]bool {
	// 如果 n 是一个明确的不会变的值，那可以考虑用数组，回避 slice 深拷贝问题
	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}

	nQueensBacktrace(board, 0)

	return nQueensResults
}

func nQueensBacktrace(board [][]bool, row int) {
	n := len(board)
	if row == n {
		// golang 的 slice 底层数据使用指针，所以这里需要进行深拷贝
		result := make([][]bool, n)
		// 并且因为是二维 slice，需要循环每一层进行拷贝
		for i := range result {
			result[i] = make([]bool, n)
			copy(result[i], board[i])
		}
		nQueensResults = append(nQueensResults, result)
		return
	}

	for col := range board[row] {
		if !isValid(board, row, col) {
			continue
		}

		board[row][col] = true
		nQueensBacktrace(board, row+1)
		board[row][col] = false
	}
}

func isValid(board [][]bool, row, col int) bool {
	n := len(board)

	for i := 0; i <= row; i++ {
		if board[i][col] {
			return false
		}
	}

	j := col + 1
	for i := row - 1; i >= 0 && j < n; i-- {
		if board[i][j] {
			return false
		}
		j++
	}

	j = col - 1
	for i := row - 1; i >= 0 && j >= 0; i-- {
		if board[i][j] {
			return false
		}
		j--
	}

	return true
}
