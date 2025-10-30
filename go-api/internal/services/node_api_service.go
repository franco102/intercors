package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type NodeAPIService struct {
	baseURL string
	client  *http.Client
}

func NewNodeAPIService() *NodeAPIService {
	baseURL := os.Getenv("NODE_API_URL")
	if baseURL == "" {
		baseURL = "http://localhost:3000"
	}

	return &NodeAPIService{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (s *NodeAPIService) authenticate() (string, error) {
	username := os.Getenv("NODE_API_USERNAME")
	password := os.Getenv("NODE_API_PASSWORD")

	if username == "" {
		username = "admin"
	}
	if password == "" {
		password = "password"
	}

	payload := map[string]string{
		"username": username,
		"password": password,
	}

	jsonPayload, _ := json.Marshal(payload)
	resp, err := s.client.Post(s.baseURL+"/login", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("login request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login failed with status %d", resp.StatusCode)
	}

	var tokenResp struct {
		Token string `json:"token"`
	}
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &tokenResp)

	if tokenResp.Token == "" {
		return "", fmt.Errorf("empty token received")
	}

	return tokenResp.Token, nil
}

func (s *NodeAPIService) SendRotatedMatrix(rotatedMatrix [][]int, originalDiagonal bool) (map[string]interface{}, error) {
	token, err := s.authenticate()
	if err != nil {
		return nil, err
	}

	payload := map[string]interface{}{
		"data":             rotatedMatrix,
		"originalDiagonal": originalDiagonal,
	}

	jsonPayload, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", s.baseURL+"/api/statistics", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	return result, nil
}
