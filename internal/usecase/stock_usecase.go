package usecase

import (
	"errors"

	delivery "github.com/sigit14ap/warehouse-service/internal/delivery/dto"
	"github.com/sigit14ap/warehouse-service/internal/domain"
	repository "github.com/sigit14ap/warehouse-service/internal/repository/mysql"
)

type StockUsecase interface {
	GetStockByWarehouse(warehouseID uint64) ([]domain.Stock, error)
	SendStock(dto delivery.SendStockDTO) (domain.Stock, error)
	TransferStock(dto delivery.TransferStockDTO) error
}

type stockUsecase struct {
	stockRepository repository.StockRepository
}

func NewStockUsecase(stockRepository repository.StockRepository) StockUsecase {
	return &stockUsecase{stockRepository}
}

func (uc *stockUsecase) GetStockByWarehouse(warehouseID uint64) ([]domain.Stock, error) {
	return uc.stockRepository.GetStockByWarehouse(warehouseID)
}

func (uc *stockUsecase) SendStock(dto delivery.SendStockDTO) (domain.Stock, error) {
	stock, _ := uc.stockRepository.GetStockByWarehouseAndProduct(dto.WarehouseID, dto.ProductID)

	if stock.ID == 0 {
		newStock := &domain.Stock{
			WarehouseID: dto.WarehouseID,
			ProductID:   dto.ProductID,
			Quantity:    dto.Quantity,
		}

		uc.stockRepository.CreateStock(newStock)
	} else {
		updatedQuantity := stock.Quantity + dto.Quantity
		uc.stockRepository.UpdateStock(dto.WarehouseID, dto.ProductID, updatedQuantity)
	}

	return uc.stockRepository.GetStockByWarehouseAndProduct(dto.WarehouseID, dto.ProductID)
}

func (uc *stockUsecase) TransferStock(dto delivery.TransferStockDTO) error {
	sourceStock, err := uc.stockRepository.GetStockByWarehouseAndProduct(dto.SourceWarehouseID, dto.ProductID)
	if err != nil {
		return err
	}
	if sourceStock.Quantity < dto.Quantity {
		return errors.New("insufficient stock in source warehouse")
	}

	updatedQuantity := sourceStock.Quantity - dto.Quantity
	err = uc.stockRepository.UpdateStock(dto.SourceWarehouseID, dto.ProductID, updatedQuantity)
	if err != nil {
		return err
	}

	destinationStock, _ := uc.stockRepository.GetStockByWarehouseAndProduct(dto.DestinationWarehouseID, dto.ProductID)

	if destinationStock.ID == 0 {
		err = uc.stockRepository.CreateStock(&domain.Stock{
			WarehouseID: dto.DestinationWarehouseID,
			ProductID:   dto.ProductID,
			Quantity:    dto.Quantity,
		})
	} else {
		updatedQuantity = destinationStock.Quantity + dto.Quantity
		err = uc.stockRepository.UpdateStock(dto.DestinationWarehouseID, dto.ProductID, updatedQuantity)
	}

	return err
}
