package main

import (
	database "crud/internal/app/configs"
	handlers "crud/internal/app/handlers"
	models "crud/internal/app/models"
	repository "crud/internal/app/repositories"
	service "crud/internal/app/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database Connection
	brandDB := database.NewInMemoryDB[models.Brand]()
	modelDB := database.NewInMemoryDB[models.Model]()

	// Initialize Repositories
	brandRepository := repository.NewCrudRepository(brandDB)
	modelRepository := repository.NewCrudRepository(modelDB)

	// Initialize Services
	brandService := service.NewBrandService(brandRepository)
	modelService := service.NewModelService(modelRepository, brandService)

	// Initialize Handlers
	brandHandler := handlers.NewBrandHandler(brandService)
	modelHandler := handlers.NewModelHandler(modelService)

	// Initialize HTTP Handlers
	router := gin.Default()
	router.POST("/brand", brandHandler.CreateBrand)
	router.GET("/brand", brandHandler.FindAll)
	router.GET("/brand/:id", brandHandler.FindById)
	router.DELETE("/brand/:id", brandHandler.DeleteById)

	router.POST("/model", modelHandler.CreateModel)
	router.GET("model", modelHandler.FindAll)
	router.GET("/model/:id", modelHandler.FindById)
	router.DELETE("/model/:id", modelHandler.DeleteById)

	router.Run(":3000")
}
