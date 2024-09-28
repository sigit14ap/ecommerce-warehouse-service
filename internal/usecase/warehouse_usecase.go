package usecase

import (
	"errors"

	"github.com/sigit14ap/warehouse-service/internal/domain"
	repository "github.com/sigit14ap/warehouse-service/internal/repository/mysql"
)

type WarehouseUsecase interface {
	GetAll() ([]domain.Warehouse, error)
	SetStatus(warehouseID uint64) error
}

type warehouseUsecase struct {
	warehouseRepository repository.WarehouseRepository
	stockRepository     repository.StockRepository
}

func NewWarehouseUsecase(warehouseRepository repository.WarehouseRepository, stockRepository repository.StockRepository) WarehouseUsecase {
	return &warehouseUsecase{
		warehouseRepository: warehouseRepository,
		stockRepository:     stockRepository,
	}
}

func (uc *warehouseUsecase) GetAll() ([]domain.Warehouse, error) {
	return uc.warehouseRepository.GetAll()
}

func (uc *warehouseUsecase) SetStatus(warehouseID uint64) error {
	warehouse, err := uc.warehouseRepository.GetByID(warehouseID)

	if err != nil {
		return err
	}

	var isActive bool

	if warehouse.IsActive {
		isActive = false

		totalStock, err := uc.stockRepository.CountTotalStockWarehouse(warehouseID)

		if err != nil {
			return err
		}

		if totalStock > 0 {
			return errors.New("cannot deactivate warehouse due to stocks available")
		}
	} else {
		isActive = true
	}

	return uc.warehouseRepository.SetStatus(warehouseID, isActive)
}
