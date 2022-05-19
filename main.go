package main

import (
	"book-api-go/book"
	"book-api-go/handler"
	"book-api-go/product"
	"book-api-go/school"
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
	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&school.School{})

	bookRepository := book.NewRepo(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	productRepository := product.NewRepo(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	schoolRepository := school.NewRepo(db)
	schoolService := school.NewService(schoolRepository)
	schoolHandler := handler.NewSchoolHandler(schoolService)

	route := gin.Default()

	v1 := route.Group("/v1")
	v1.GET("/books", bookHandler.GetBooksHandler)
	v1.GET("/books/:id", bookHandler.GetBookHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)
	v1.PATCH("/books/:id", bookHandler.UpdateBookHandler)

	v2 := route.Group("/v2")
	v2.GET("/products", productHandler.GetProductsHeader)
	v2.GET("/products/:id", productHandler.GetProductHandler)
	v2.POST("/products", productHandler.PostProductsHandler)
	v2.DELETE("/products/:id", productHandler.DeleteProductHandler)
	v2.PATCH("/products/:id", productHandler.UpdateProductHandler)

	route.GET("/", schoolHandler.GetSchoolsHeader)
	route.GET("/:id", schoolHandler.GetSchoolHandler)
	route.POST("/", schoolHandler.PostSchoolHandler)
	route.DELETE("/:id", schoolHandler.DeleteSchoolHandler)
	route.PATCH("/:id", schoolHandler.UpdateSchoolHandler)

	route.Run("localhost:8080")

}
