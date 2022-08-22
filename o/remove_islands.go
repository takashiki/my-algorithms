package o

// https://www.youtube.com/watch?v=4tYoVx0QoN0
func RemoveIslands(input [][]bool) [][]bool {
	rowNum := len(input)
	colNum := len(input[0])

	output := make([][]bool, rowNum)
	for i := range output {
		output[i] = make([]bool, colNum)
	}

	// there may be duplicate works in the central points can be enhanced, but it's not a big deal.
	for row := 0; row <= rowNum/2; row++ {
		for col := 0; col < colNum/2; col++ {
			if isConnectedToBorder(input, output, row, col) {
				output[row][col] = true
			}

			if isConnectedToBorder(input, output, row, colNum-1-col) {
				output[row][colNum-1-col] = true
			}

			if isConnectedToBorder(input, output, rowNum-1-row, col) {
				output[rowNum-1-row][col] = true
			}

			if isConnectedToBorder(input, output, rowNum-1-row, colNum-1-col) {
				output[rowNum-1-row][colNum-1-col] = true
			}
		}
	}

	return output
}

func isConnectedToBorder(input [][]bool, output [][]bool, row, col int) bool {
	// if this point is 0 in the input, it's not connected to the border.
	if !input[row][col] {
		return false
	}

	// Now we know that the current point is 1.
	// If the current point is at the border, return true.
	if row == 0 || col == 0 || row == len(input)-1 || col == len(input[0])-1 {
		return true
	}

	// If the left or right or up or down point is connected to the border, return true.
	// ENH: we can check the side of the current point and not to read uninitialed adjacent points.
	if output[row][col-1] || output[row][col+1] || output[row-1][col] || output[row+1][col] {
		return true
	}

	// Otherwise, return false.
	return false
}
