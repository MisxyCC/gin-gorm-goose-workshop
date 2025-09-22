package main

import (
	"net/http"

	"example.com/gin-gorm-goose/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	dsn := "host=localhost user=ne1User password=mysecretpassword dbname=sampleDB port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to databse!")
	}

	router := gin.Default()

	router.POST("/products", func(ctx *gin.Context) {
		var product models.Product
		if err := ctx.ShouldBindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := DB.Create(&product)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, product)
	})

	router.GET("/products", func(ctx *gin.Context) {
		var products []models.Product
		result := DB.Find(&products)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
			return
		}
		ctx.JSON(http.StatusOK, products)
	})
	router.Run(":8080")
}