package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type Product struct {
	ID          int64   `json:"id" validate:"required"`
	Name        string  `json:"name" validate:"required,max=30"`
	Description string  `json:"description" validate:"max=150"`
	Price       float64 `json:"price" validate:"required,min=1,max=100"`
	Currency    string  `json:"currency" validate:"required,iso4217"`
}

func RegisterProduct(gr *gin.RouterGroup) {
	v1 := gr.Group("/v1")
	v1.POST("/products", func(c *gin.Context) {
		var product Product

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"status":  http.StatusText(http.StatusInternalServerError),
				"message": "error parsing body",
			})
			return
		}

		if err := validate.Struct(product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"status":  http.StatusText(http.StatusBadRequest),
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, product)
	})
}
