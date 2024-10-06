package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseHost      string
	DatabasePort      string
	DatabaseUser      string
	DatabasePassword  string
	DatabaseName      string
	AppSecret         string
	ShopServiceUrl    string
	ProductServiceUrl string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		DatabaseHost:      os.Getenv("DATABASE_HOST"),
		DatabasePort:      os.Getenv("DATABASE_PORT"),
		DatabaseUser:      os.Getenv("DATABASE_USER"),
		DatabasePassword:  os.Getenv("DATABASE_PASSWORD"),
		DatabaseName:      os.Getenv("DATABASE_NAME"),
		AppSecret:         os.Getenv("APP_SECRET"),
		ShopServiceUrl:    os.Getenv("SHOP_SERVICE_BASE_URL"),
		ProductServiceUrl: os.Getenv("PRODUCT_SERVICE_BASE_URL"),
	}

	return config
}
