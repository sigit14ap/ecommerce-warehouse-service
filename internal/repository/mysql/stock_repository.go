package repository

import (
	"context"

	"github.com/sigit14ap/warehouse-service/internal/domain"
	"gorm.io/gorm"
)

type StockRepository interface {
	GetStockByWarehouseAndProduct(ctx context.Context, warehouseID, productID uint64) (domain.Stock, error)
	UpdateStock(ctx context.Context, warehouseID, productID uint64, quantity int) error
	CreateStock(ctx context.Context, stock *domain.Stock) error
	GetStockByWarehouse(ctx context.Context, warehouseID uint64) ([]domain.Stock, error)
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockRepository{db}
}

func (repository *stockRepository) GetStockByWarehouseAndProduct(ctx context.Context, warehouseID, productID uint64) (domain.Stock, error) {
	var stock domain.Stock
	err := repository.db.Where("warehouse_id = ? AND product_id = ?", warehouseID, productID).First(&stock).Error
	return stock, err
}

func (repository *stockRepository) UpdateStock(ctx context.Context, warehouseID, productID uint64, quantity int) error {
	return repository.db.Model(&domain.Stock{}).Where("warehouse_id = ? AND product_id = ?", warehouseID, productID).Update("quantity", quantity).Error
}

func (repository *stockRepository) CreateStock(ctx context.Context, stock *domain.Stock) error {
	return repository.db.Create(stock).Error
}

func (repository *stockRepository) GetStockByWarehouse(ctx context.Context, warehouseID uint64) ([]domain.Stock, error) {
	var stock []domain.Stock
	err := repository.db.Where("warehouse_id = ?", warehouseID).Find(&stock).Error
	return stock, err
}
