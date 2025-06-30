package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fathimasithara01/productService/models"
	"github.com/fathimasithara01/productService/views"
)

type ProductController struct{ DB *gorm.DB }

func New(db *gorm.DB) *ProductController { return &ProductController{DB: db} }

// POST /products
func (pc *ProductController) Create(c *gin.Context) {
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		views.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := pc.DB.Create(&p).Error; err != nil {
		views.Error(c, http.StatusInternalServerError, "create failed")
		return
	}
	views.Created(c, p)
}

// GET /products/:id
func (pc *ProductController) Get(c *gin.Context) {
	var p models.Product
	if err := pc.DB.First(&p, c.Param("id")).Error; err != nil {
		views.Error(c, http.StatusNotFound, "not found")
		return
	}
	views.OK(c, p)
}

// GET /products
func (pc *ProductController) List(c *gin.Context) {
	var list []models.Product
	pc.DB.Find(&list)
	views.OK(c, list)
}

// PUT /products/:id
func (pc *ProductController) Update(c *gin.Context) {
	var p models.Product
	if err := pc.DB.First(&p, c.Param("id")).Error; err != nil {
		views.Error(c, http.StatusNotFound, "not found")
		return
	}
	if err := c.ShouldBindJSON(&p); err != nil {
		views.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := pc.DB.Save(&p).Error; err != nil {
		views.Error(c, http.StatusInternalServerError, "update failed")
		return
	}
	views.OK(c, p)
}

// DELETE /products/:id
func (pc *ProductController) Delete(c *gin.Context) {
	if err := pc.DB.Delete(&models.Product{}, c.Param("id")).Error; err != nil {
		views.Error(c, http.StatusInternalServerError, "delete failed")
		return
	}
	views.OK(c, gin.H{"message": "deleted"})
}
