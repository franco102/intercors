package main

import (
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name: "3x3 matrix",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "2x3 matrix",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expected: [][]int{
				{4, 1},
				{5, 2},
				{6, 3},
			},
		},
		{
			name: "1x1 matrix",
			input: [][]int{
				{5},
			},
			expected: [][]int{
				{5},
			},
		},
		{
			name:     "empty matrix",
			input:    [][]int{},
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RotateMatrix(tt.input)

			if len(result) != len(tt.expected) {
				t.Errorf("expected %d rows, got %d", len(tt.expected), len(result))
			}

			for i := range result {
				for j := range result[i] {
					if result[i][j] != tt.expected[i][j] {
						t.Errorf("at [%d][%d]: expected %d, got %d", i, j, tt.expected[i][j], result[i][j])
					}
				}
			}
		})
	}
}
