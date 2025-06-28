package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fathimasithara01/advance/controllers"
	"github.com/fathimasithara01/advance/middleware"
)

func Register(r *gin.Engine, db *gorm.DB) {
	uc := controllers.New(db)
	r.POST("/register", uc.Register)
	r.POST("/login", uc.Login)

	r.GET("/me", middleware.AuthRequired(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Authenticated successfully"})
	})
}
