package dto

import "time"

type ApiResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ProductDetailResponse struct {
	ID        uint64    `json:"id"`
	ShopID    string    `json:"shop_id"`
	Name      string    `json:"name"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
