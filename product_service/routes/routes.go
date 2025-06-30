package routes

import (
	// "product-service/middlewares" // enable for protected routes

	"github.com/fathimasithara01/productService/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, db *gorm.DB) {
	pc := controllers.New(db)

	r.GET("/products", pc.List)
	r.GET("/products/:id", pc.Get)

	// For auth-protected create/update/delete, wrap with middlewares.AuthRequired()
	// auth := middlewares.AuthRequired()
	r.POST("/products", pc.Create)
	r.PUT("/products/:id", pc.Update)
	r.DELETE("/products/:id", pc.Delete)
}
