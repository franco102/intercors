package matrix

// IsDiagonal checks if a matrix is diagonal
func IsDiagonal(matrix [][]int) bool {
	if len(matrix) == 0 {
		return false
	}

	rows := len(matrix)
	cols := len(matrix[0])

	if rows != cols {
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i != j && matrix[i][j] != 0 {
				return false
			}
		}
	}

	return true
}

// IsValid checks if matrix is valid (non-empty and rectangular)
func IsValid(matrix [][]int) bool {
	if len(matrix) == 0 {
		return false
	}

	cols := len(matrix[0])
	for i := 1; i < len(matrix); i++ {
		if len(matrix[i]) != cols {
			return false
		}
	}

	return true
}
