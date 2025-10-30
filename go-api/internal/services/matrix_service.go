package services

import (
	"fmt"
	"go-api/pkg/matrix"
)

type MatrixService struct {
	nodeAPIService *NodeAPIService
}

func NewMatrixService(nodeAPIService *NodeAPIService) *MatrixService {
	return &MatrixService{
		nodeAPIService: nodeAPIService,
	}
}

func (s *MatrixService) RotateMatrix(data [][]int) ([][]int, map[string]interface{}, error) {
	// Validate matrix
	if !matrix.IsValid(data) {
		return nil, nil, fmt.Errorf("invalid matrix: must be non-empty and rectangular")
	}

	// Check if original matrix is diagonal
	isOriginalDiagonal := matrix.IsDiagonal(data)

	// Rotate the matrix
	rotatedMatrix := matrix.Rotate(data)

	// Send to Node.js API for statistics
	statistics, err := s.nodeAPIService.SendRotatedMatrix(rotatedMatrix, isOriginalDiagonal)
	if err != nil {
		return nil, nil, fmt.Errorf("error sending to Node.js API: %v", err)
	}

	return rotatedMatrix, statistics, nil
}
