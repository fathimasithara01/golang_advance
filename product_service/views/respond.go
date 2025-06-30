package views

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type envelope struct {
    Status  string      `json:"status"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
}

func OK(c *gin.Context, d interface{})      { c.JSON(http.StatusOK, envelope{"success", "", d}) }
func Created(c *gin.Context, d interface{}) { c.JSON(http.StatusCreated, envelope{"success", "", d}) }
func Error(c *gin.Context, code int, msg string) {
    c.JSON(code, envelope{"error", msg, nil})
}
