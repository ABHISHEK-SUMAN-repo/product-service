package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-service/dto"
	"product-service/service"
)

func CreateProducts(c *gin.Context) {
	var productDTO []dto.ProductDTO

	if err := c.Bind(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := service.CreateProducts(productDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": createdProduct, "status": true, "code": http.StatusOK})
}

func GetProductsByIds(c *gin.Context) {
	var ids []string

	if err := c.Bind(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := service.GetProductsByIds(ids)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, response)
}