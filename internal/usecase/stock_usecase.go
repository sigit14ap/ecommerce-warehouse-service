package usecase

import (
	"context"
	"errors"

	delivery "github.com/sigit14ap/warehouse-service/internal/delivery/dto"
	"github.com/sigit14ap/warehouse-service/internal/domain"
	repository "github.com/sigit14ap/warehouse-service/internal/repository/mysql"
)

type StockUsecase interface {
	GetWarehouseStock(ctx context.Context, warehouseID uint64) ([]domain.Stock, error)
	TransferStock(ctx context.Context, dto delivery.TransferProductDTO) error
}

type stockUsecase struct {
	stockRepository repository.StockRepository
}

func NewStockUsecase(stockRepository repository.StockRepository) StockUsecase {
	return &stockUsecase{stockRepository}
}

func (uc *stockUsecase) GetWarehouseStock(ctx context.Context, warehouseID uint64) ([]domain.Stock, error) {
	return uc.stockRepository.GetStockByWarehouse(ctx, warehouseID)
}

func (uc *stockUsecase) TransferStock(ctx context.Context, dto delivery.TransferProductDTO) error {
	sourceStock, err := uc.stockRepository.GetStockByWarehouseAndProduct(ctx, dto.SourceWarehouseID, dto.ProductID)
	if err != nil {
		return err
	}
	if sourceStock.Quantity < dto.Quantity {
		return errors.New("insufficient stock in source warehouse")
	}

	err = uc.stockRepository.UpdateStock(ctx, dto.SourceWarehouseID, dto.ProductID, sourceStock.Quantity-dto.Quantity)
	if err != nil {
		return err
	}

	destinationStock, _ := uc.stockRepository.GetStockByWarehouseAndProduct(ctx, dto.DestinationWarehouseID, dto.ProductID)
	if destinationStock.ID == 0 {
		err = uc.stockRepository.CreateStock(ctx, &domain.Stock{
			WarehouseID: dto.DestinationWarehouseID,
			ProductID:   dto.ProductID,
			Quantity:    dto.Quantity,
		})
	} else {
		err = uc.stockRepository.UpdateStock(ctx, dto.DestinationWarehouseID, dto.ProductID, destinationStock.Quantity+dto.Quantity)
	}

	return err
}
