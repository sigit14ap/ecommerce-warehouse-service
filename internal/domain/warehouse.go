package domain

import "time"

type Warehouse struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255" json:"name"`
	Location  string    `gorm:"size:255" json:"location"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Stock struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	WarehouseID uint64    `json:"warehouse_id"`
	ProductID   uint64    `json:"product_id"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
