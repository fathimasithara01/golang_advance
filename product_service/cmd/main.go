package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/fathimasithara01/productService/config"
	"github.com/fathimasithara01/productService/models"
	"github.com/fathimasithara01/productService/routes"
	"github.com/fathimasithara01/productService/utils"
)

func main() {
	config.Load()

	db, err := gorm.Open(postgres.Open(viper.GetString("DB_URL")), &gorm.Config{})
	if err != nil {
		utils.Error("db connect", err)
		os.Exit(1)
	}
	_ = db.AutoMigrate(&models.Product{})

	r := gin.Default()
	routes.Register(r, db)

	utils.Info("product-service on " + viper.GetString("PORT"))
	if err := r.Run(":" + viper.GetString("PORT")); err != nil {
		utils.Error("server exit", err)
	}
}
