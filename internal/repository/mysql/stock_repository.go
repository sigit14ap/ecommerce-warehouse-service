package repository

import (
	"github.com/sigit14ap/warehouse-service/internal/domain"
	"gorm.io/gorm"
)

type StockRepository interface {
	GetStockByWarehouseAndProduct(warehouseID uint64, productID uint64) (domain.Stock, error)
	UpdateStock(warehouseID uint64, productID uint64, quantity int) error
	CreateStock(stock *domain.Stock) error
	GetStockByWarehouse(warehouseID uint64) ([]domain.Stock, error)
	CountTotalStockWarehouse(warehouseID uint64) (int64, error)
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockRepository{db}
}

func (repository *stockRepository) GetStockByWarehouseAndProduct(warehouseID, productID uint64) (domain.Stock, error) {
	var stock domain.Stock
	err := repository.db.Where("warehouse_id = ? AND product_id = ?", warehouseID, productID).First(&stock).Error
	return stock, err
}

func (repository *stockRepository) UpdateStock(warehouseID, productID uint64, quantity int) error {
	return repository.db.Model(&domain.Stock{}).Where("warehouse_id = ? AND product_id = ?", warehouseID, productID).Update("quantity", quantity).Error
}

func (repository *stockRepository) CreateStock(stock *domain.Stock) error {
	return repository.db.Create(stock).Error
}

func (repository *stockRepository) GetStockByWarehouse(warehouseID uint64) ([]domain.Stock, error) {
	var stock []domain.Stock
	err := repository.db.Where("warehouse_id = ?", warehouseID).Find(&stock).Error
	return stock, err
}

func (repository *stockRepository) CountTotalStockWarehouse(warehouseID uint64) (int64, error) {
	var totalStock int64

	err := repository.db.Table("stocks").Select("IFNULL(SUM(quantity), 0)").Where("warehouse_id = ?", warehouseID).Scan(&totalStock).Error

	if err != nil {
		return 0, err
	}

	return totalStock, nil
}
