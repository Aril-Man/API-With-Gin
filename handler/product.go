package handler

import (
	"book-api-go/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{
		productService: productService,
	}
}

func (h *productHandler) GetProductsHeader(c *gin.Context) {
	products, err := h.productService.GetProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var responseProduct []product.ProductResponse

	for _, b := range products {
		productResponse := convertProductRequestToProduct(b)

		responseProduct = append(responseProduct, productResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": responseProduct,
	})
}

func (h *productHandler) GetProductHandler(c *gin.Context) {
	idString := c.Param("id")

	id, _ := strconv.Atoi(idString)

	b, err := h.productService.GetProduct(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	productResponse := convertProductRequestToProduct(b)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    productResponse,
		"message": "success mengambil data book",
	})
}

func (h *productHandler) PostProductsHandler(c *gin.Context) {
	var productRequest product.ProductRequest

	err := c.ShouldBindJSON(&productRequest)

	if err != nil {
		errorMessage := []string{}

		for _, err := range err.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, err.Field()+": "+err.Tag())
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessage,
		})
		return
	}

	product, err := h.productService.CreateProduct(productRequest)

	productResponse := convertProductRequestToProduct(product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productResponse,
	})
}

func (h *productHandler) DeleteProductHandler(c *gin.Context) {
	idString := c.Param("id")

	id, _ := strconv.Atoi(idString)

	_, err := h.productService.DeleteProduct(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
	})
}

func (h *productHandler) UpdateProductHandler(c *gin.Context) {
	var productRequest product.ProductRequest

	err := c.ShouldBindJSON(&productRequest)

	if err != nil {
		errorMessage := []string{}

		for _, err := range err.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, err.Field()+": "+err.Tag())
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessage,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.productService.UpdateProduct(id, productRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	productResponse := convertProductRequestToProduct(b)

	c.JSON(http.StatusOK, gin.H{
		"data": productResponse,
	})
}

func convertProductRequestToProduct(b product.Product) product.ProductResponse {
	return product.ProductResponse{
		ID:      b.ID,
		Product: b.Product,
		Price:   b.Price,
	}
}
