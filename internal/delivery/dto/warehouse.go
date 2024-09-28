package delivery

type ToggleWarehouseStatusDTO struct {
	IsActive bool `json:"is_active" validate:"required"`
}

type TransferProductDTO struct {
	SourceWarehouseID      uint64 `json:"source_warehouse_id" validate:"required,gt=0"`
	DestinationWarehouseID uint64 `json:"destination_warehouse_id" validate:"required,gt=0"`
	ProductID              uint64 `json:"product_id" validate:"required,gt=0"`
	Quantity               int    `json:"quantity" validate:"required,min=1"`
}
