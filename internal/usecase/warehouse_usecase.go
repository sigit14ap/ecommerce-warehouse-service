package usecase

import (
	"github.com/sigit14ap/warehouse-service/internal/domain"
	repository "github.com/sigit14ap/warehouse-service/internal/repository/mysql"
)

type WarehouseUsecase interface {
	GetAll() ([]domain.Warehouse, error)
	SetStatus(warehouseID uint64) error
}

type warehouseUsecase struct {
	warehouseRepo repository.WarehouseRepository
}

func NewWarehouseUsecase(warehouseRepository repository.WarehouseRepository) WarehouseUsecase {
	return &warehouseUsecase{warehouseRepository}
}

func (uc *warehouseUsecase) GetAll() ([]domain.Warehouse, error) {
	return uc.warehouseRepo.GetAll()
}

func (uc *warehouseUsecase) SetStatus(warehouseID uint64) error {
	warehouse, err := uc.warehouseRepo.GetByID(warehouseID)

	if err != nil {
		return err
	}

	var isActive bool

	if warehouse.IsActive {
		isActive = false
	} else {
		isActive = true
	}

	return uc.warehouseRepo.SetStatus(warehouseID, isActive)
}
