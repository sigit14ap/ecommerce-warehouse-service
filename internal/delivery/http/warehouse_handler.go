package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sigit14ap/warehouse-service/helpers"
	"github.com/sigit14ap/warehouse-service/internal/usecase"
)

type WarehouseHandler struct {
	warehouseUsecase usecase.WarehouseUsecase
}

func NewWarehouseHandler(warehouseUsecase usecase.WarehouseUsecase) *WarehouseHandler {
	return &WarehouseHandler{warehouseUsecase: warehouseUsecase}
}

func (handler *WarehouseHandler) GetAll(context *gin.Context) {
	warehouses, err := handler.warehouseUsecase.GetAll()
	if err != nil {
		helpers.ErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.SuccessResponse(context, warehouses)
}

func (handler *WarehouseHandler) SetStatus(context *gin.Context) {
	id := context.Param("id")

	warehouseID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Warehouse ID")
		return
	}

	if err := handler.warehouseUsecase.SetStatus(warehouseID); err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(context, nil)
}
