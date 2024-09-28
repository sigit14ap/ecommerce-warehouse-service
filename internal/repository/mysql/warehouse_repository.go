package repository

import (
	"github.com/sigit14ap/warehouse-service/internal/domain"
	"gorm.io/gorm"
)

type WarehouseRepository interface {
	GetAll() ([]domain.Warehouse, error)
	GetByID(id uint64) (*domain.Warehouse, error)
	SetStatus(warehouseID uint64, isActive bool) error
}

type warehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) WarehouseRepository {
	return &warehouseRepository{db}
}

func (repository *warehouseRepository) GetAll() ([]domain.Warehouse, error) {
	var warehouses []domain.Warehouse
	if err := repository.db.Find(&warehouses).Error; err != nil {
		return nil, err
	}

	return warehouses, nil
}

func (repository *warehouseRepository) GetByID(id uint64) (*domain.Warehouse, error) {
	var warehouse domain.Warehouse
	if err := repository.db.Where("id = ?", id).First(&warehouse).Error; err != nil {
		return nil, err
	}

	return &warehouse, nil
}

func (repository *warehouseRepository) SetStatus(warehouseID uint64, isActive bool) error {
	return repository.db.Model(&domain.Warehouse{}).Where("id = ?", warehouseID).Update("is_active", isActive).Error
}
