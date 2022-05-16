package book

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title" `
	Price       int    `json:"price" `
	Description string `json:"description"`
	Rating      int    `json:"rating" `
}
