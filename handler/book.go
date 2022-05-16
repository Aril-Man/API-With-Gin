package handler

import (
	"book-api-go/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{
		bookService: bookService,
	}
}

func (h *bookHandler) GetBooksHandler(c *gin.Context) {
	books, err := h.bookService.GetBooks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertBookRequestToBook(b)

		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBookHandler(c *gin.Context) {
	idString := c.Param("id")

	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.GetBook(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	bookResponse := convertBookRequestToBook(b)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    bookResponse,
		"message": "success mengambil data book",
	})
}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

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

	book, err := h.bookService.CreateBook(bookRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) DeleteBookHandler(c *gin.Context) {
	idString := c.Param("id")

	id, _ := strconv.Atoi(idString)

	_, err := h.bookService.DeleteBook(id)

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

func (h *bookHandler) UpdateBookHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

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

	b, err := h.bookService.UpdateBook(id, bookRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	bookResponse := convertBookRequestToBook(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func convertBookRequestToBook(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
	}
}
