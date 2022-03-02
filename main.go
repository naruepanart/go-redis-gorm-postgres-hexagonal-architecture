package main

import (
	"goredis/handlers"
	"goredis/repositories"
	"goredis/services"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redisClient := initRedis()
	_ = redisClient

	productRepo := repositories.NewProductRepositoryDB(db)
	productService := services.NewCatalogServiceRedis(productRepo, redisClient)
	productHandler := handlers.NewCatalogHandler(productService)

	app := fiber.New()
	app.Get("/products", productHandler.GetProducts)
	app.Listen(":3101")

}

func initDatabase() *gorm.DB {
	dial := "host=localhost user=postgresql password=U9Ni8JJp3LnJYBCR dbname=edugorm port=5432 TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dial), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "password123",
		ReadTimeout: 10 * time.Minute,
	})
}
