package pascaltriangle

/**
[0]
[0, 0]
[0, 0+1, 0]
[0, 0+(0+1), (0+1)+0, 0]
**/

func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	if numRows >= 1 {
		result[0] = []int{1}
	}

	if numRows >= 2 {
		result[1] = []int{1, 1}
	}

	for i := 2; i < numRows; i++ {
		row := make([]int, i+1)
		lastRow := result[i-1]

		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else {
				row[j] = lastRow[j] + lastRow[j-1]
			}
		}

		result[i] = row
	}

	return result
}
