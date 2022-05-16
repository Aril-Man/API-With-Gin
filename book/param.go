package book

type BookRequest struct {
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}
