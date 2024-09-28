package router

import (
	delivery "github.com/sigit14ap/warehouse-service/internal/delivery/http"
	"github.com/sigit14ap/warehouse-service/internal/middleware"
	"github.com/sigit14ap/warehouse-service/internal/services"

	"github.com/gin-gonic/gin"
)

func NewRouter(warehouseHandler *delivery.WarehouseHandler, shopClient *services.ShopClient) *gin.Engine {
	router := gin.New()
	v1 := router.Group("/api/v1")
	v1.Use(middleware.ServiceMiddleware())

	warehouse := v1.Group("warehouses")
	{
		warehouse.GET("/", warehouseHandler.GetAll)
		warehouse.PATCH("/:id/status", warehouseHandler.SetStatus)
	}

	return router
}
