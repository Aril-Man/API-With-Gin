package product

type ProductRequest struct {
	Product string `json:"product" binding:"required"`
	Price   int    `json:"price" binding:"required,number"`
}
