package main

import (
	"fmt"
	"os"

	"github.com/sigit14ap/warehouse-service/config"
	"github.com/sigit14ap/warehouse-service/helpers"
	delivery "github.com/sigit14ap/warehouse-service/internal/delivery/http"
	"github.com/sigit14ap/warehouse-service/internal/domain"
	repository "github.com/sigit14ap/warehouse-service/internal/repository/mysql"
	"github.com/sigit14ap/warehouse-service/internal/router"
	"github.com/sigit14ap/warehouse-service/internal/services"
	"github.com/sigit14ap/warehouse-service/internal/usecase"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	log := helpers.InitializeLogs()

	if log == nil {
		log.Fatal("Logger failed to started")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&domain.Warehouse{}, &domain.Stock{})
	if err != nil {
		log.Fatalf("failed to auto-migrate Warehouse model: %v", err)
	}

	seedWarehouseData(db, log)

	shopService := services.NewShopService(cfg.ShopServiceUrl, cfg.AppSecret)
	productService := services.NewProductService(cfg.ProductServiceUrl, cfg.AppSecret)

	stockRepo := repository.NewStockRepository(db)
	stockUsecase := usecase.NewStockUsecase(stockRepo)
	stockHandler := delivery.NewStockHandler(stockUsecase, productService)

	warehouseRepo := repository.NewWarehouseRepository(db)
	warehouseUsecase := usecase.NewWarehouseUsecase(warehouseRepo, stockRepo)
	warehouseHandler := delivery.NewWarehouseHandler(warehouseUsecase)

	router := router.NewRouter(warehouseHandler, stockHandler, shopService)

	log.Info(router.Run(":" + os.Getenv("APP_PORT")))
}

func seedWarehouseData(db *gorm.DB, log *logrus.Logger) {
	var count int64
	db.Model(&domain.Warehouse{}).Count(&count)
	if count > 0 {
		log.Info("Warehouses already seeded, skipping...")
		return
	}

	warehouses := []domain.Warehouse{
		{Name: "Jakarta Warehouse", Location: "Jakarta", IsActive: true},
		{Name: "Bandung Warehouse", Location: "Bandung", IsActive: true},
		{Name: "Surabaya Warehouse", Location: "Surabaya", IsActive: true},
	}

	err := db.Create(&warehouses).Error
	if err != nil {
		log.Fatal("Failed to seed warehouse data: ", err)
	} else {
		log.Info("Successfully seeded warehouse data.")
	}
}
