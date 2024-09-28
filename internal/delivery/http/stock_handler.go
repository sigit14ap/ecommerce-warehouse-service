package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/sigit14ap/warehouse-service/helpers"
	"github.com/sigit14ap/warehouse-service/internal/delivery/dto"
	"github.com/sigit14ap/warehouse-service/internal/services"
	"github.com/sigit14ap/warehouse-service/internal/usecase"
)

type StockHandler struct {
	stockUsecase   usecase.StockUsecase
	productService *services.ProductService
}

var validate *validator.Validate

func NewStockHandler(stockUsecase usecase.StockUsecase, productService *services.ProductService) *StockHandler {
	return &StockHandler{
		stockUsecase:   stockUsecase,
		productService: productService,
	}
}

func (handler *StockHandler) GetStockByWarehouse(context *gin.Context) {
	id := context.Param("id")

	warehouseID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Warehouse ID")
		return
	}

	stocks, err := handler.stockUsecase.GetStockByWarehouse(warehouseID)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(context, stocks)
}

func (handler *StockHandler) SendStock(context *gin.Context) {
	validate = validator.New()

	var sendStockRequest dto.SendStockDTO

	err := context.BindJSON(&sendStockRequest)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = validate.Struct(sendStockRequest)
	if err != nil {
		helpers.ErrorValidationResponse(context, err)
		return
	}

	_, err = handler.productService.ProductDetail(context, sendStockRequest.ProductID)

	if err != nil {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "Unauthorized")
		return
	}

	stock, err := handler.stockUsecase.SendStock(sendStockRequest)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(context, stock)
}

func (handler *StockHandler) TransferStock(context *gin.Context) {
	validate = validator.New()

	var transferRequest dto.TransferStockDTO

	err := context.BindJSON(&transferRequest)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = validate.Struct(transferRequest)
	if err != nil {
		helpers.ErrorValidationResponse(context, err)
		return
	}

	_, err = handler.productService.ProductDetail(context, transferRequest.ProductID)

	if err != nil {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "Unauthorized")
		return
	}

	err = handler.stockUsecase.TransferStock(transferRequest)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(context, nil)
}
