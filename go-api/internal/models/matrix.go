package models

// Matrix represents a 2D array of integers
type Matrix struct {
	Data [][]int `json:"data"`
}

// RotatedMatrix represents the rotated matrix to send to Node.js API
type RotatedMatrix struct {
	Data             [][]int `json:"data"`
	OriginalDiagonal bool    `json:"originalDiagonal"`
}

// RotateResponse represents the API response
type RotateResponse struct {
	RotatedMatrix [][]int                `json:"rotatedMatrix"`
	Statistics    map[string]interface{} `json:"statistics"`
}
