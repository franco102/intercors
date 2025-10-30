package matrix

// Rotate rotates a matrix 90 degrees clockwise
func Rotate(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}

	rows := len(matrix)
	cols := len(matrix[0])

	rotated := make([][]int, cols)
	for i := range rotated {
		rotated[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = matrix[i][j]
		}
	}

	return rotated
}
