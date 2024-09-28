package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sigit14ap/warehouse-service/internal/delivery/dto"
)

type ProductService struct {
	BaseURL      string
	ServiceToken string
}

func NewProductService(baseURL, token string) *ProductService {
	return &ProductService{
		BaseURL:      baseURL,
		ServiceToken: token,
	}
}

func (c *ProductService) CallProductService(method, endpoint string, authToken string, payload interface{}) (*http.Response, error) {
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

func (c *ProductService) ProductDetail(context *gin.Context, productID uint64) (*dto.ProductDetailResponse, error) {
	token := context.GetHeader("Authorization")
	if token == "" {
		return nil, errors.New("Authorization required")
	}

	url := fmt.Sprintf("%s%d", "api/v1/shop/products/", productID)
	response, err := c.CallProductService("GET", url, token, nil)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var apiResponse dto.ApiResponse
	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	productData, ok := apiResponse.Data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected data format")
	}

	productDetail, err := json.Marshal(productData["shop"])
	if err != nil {
		return nil, fmt.Errorf("failed to marshal shop data: %w", err)
	}

	var shop dto.ProductDetailResponse
	if err := json.Unmarshal(productDetail, &shop); err != nil {
		return nil, fmt.Errorf("failed to unmarshal shop data: %w", err)
	}

	return &shop, nil
}
