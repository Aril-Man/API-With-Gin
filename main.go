package main

import (
	"book-api-go/book"
	"book-api-go/handler"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/api-go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepo(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	route := gin.Default()

	v1 := route.Group("/v1")
	v1.GET("/books", bookHandler.GetBooksHandler)
	v1.GET("/books/:id", bookHandler.GetBookHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)
	v1.PATCH("/books/:id", bookHandler.UpdateBookHandler)

	route.Run(":8080")

}
