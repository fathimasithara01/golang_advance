package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/fathimasithara01/advance/models"
	"github.com/fathimasithara01/advance/utils"
	"github.com/fathimasithara01/advance/views"
)

type UserController struct{ DB *gorm.DB }

func New(db *gorm.DB) *UserController { return &UserController{DB: db} }

func (uc *UserController) Register(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		views.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hash)
	if err := uc.DB.Create(&u).Error; err != nil {
		views.Error(c, http.StatusConflict, "email already exists")
		return
	}
	views.Created(c, gin.H{"id": u.ID, "email": u.Email})
}

func (uc *UserController) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		views.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var u models.User
	if err := uc.DB.Where("email=?", req.Email).First(&u).Error; err != nil {
		views.Error(c, http.StatusUnauthorized, "invalid credentials")
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)) != nil {
		views.Error(c, http.StatusUnauthorized, "invalid credentials")
		return
	}

	token, _ := utils.GenerateToken(u.ID)
	views.OK(c, gin.H{"token": token})
}
