package dto

type ToggleWarehouseStatusDTO struct {
	IsActive bool `json:"is_active" validate:"required"`
}

type SendStockDTO struct {
	WarehouseID uint64 `json:"warehouse_id" validate:"required"`
	ProductID   uint64 `json:"product_id" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required,min=1"`
}

type TransferStockDTO struct {
	SourceWarehouseID      uint64 `json:"source_warehouse_id" validate:"required"`
	DestinationWarehouseID uint64 `json:"destination_warehouse_id" validate:"required"`
	ProductID              uint64 `json:"product_id" validate:"required,gt=0"`
	Quantity               int    `json:"quantity" validate:"required,min=1"`
}
