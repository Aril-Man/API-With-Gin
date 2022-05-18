package product

type ProductResponse struct {
	ID      int    `json:"id"`
	Product string `json:"product"`
	Price   int    `json:"price"`
}
