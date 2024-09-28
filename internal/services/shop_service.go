package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ShopClient struct {
	BaseURL      string
	ServiceToken string
}

func NewShopClient(baseURL, token string) *ShopClient {
	return &ShopClient{
		BaseURL:      baseURL,
		ServiceToken: token,
	}
}

type ApiResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ShopDetailResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *ShopClient) CallShopService(method, endpoint string, authToken string, payload interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)

	var jsonData []byte
	var err error
	if payload != nil {
		jsonData, err = json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %w", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)
	req.Header.Set("X-Service-Token", c.ServiceToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call shop service: %w", err)
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return resp, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("unauthorized access")
	case http.StatusNotFound:
		return nil, fmt.Errorf("resource not found")
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("internal server error")
	default:
		return nil, fmt.Errorf("unexpected status code %d received from shop service", resp.StatusCode)
	}
}

func (c *ShopClient) ShopDetail(authToken string) (*ShopDetailResponse, error) {
	response, err := c.CallShopService("GET", "api/v1/shop/me", authToken, nil)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var apiResponse ApiResponse
	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	shopData, ok := apiResponse.Data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected data format")
	}

	shopDetail, err := json.Marshal(shopData["shop"])
	if err != nil {
		return nil, fmt.Errorf("failed to marshal shop data: %w", err)
	}

	var shop ShopDetailResponse
	if err := json.Unmarshal(shopDetail, &shop); err != nil {
		return nil, fmt.Errorf("failed to unmarshal shop data: %w", err)
	}

	return &shop, nil
}
