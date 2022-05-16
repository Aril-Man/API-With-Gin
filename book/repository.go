package book

import "gorm.io/gorm"

type Repository interface {
	GetBooks() ([]Book, error)
	GetBook(id int) (Book, error)
	CreateBook(book Book) (Book, error)
	DeleteBook(id int) (Book, error)
	UpdateBook(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetBooks() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error

	if err != nil {
		return nil, err
	}

	return books, err
}

func (r *repository) GetBook(id int) (Book, error) {
	var book Book
	err := r.db.First(&book, id).Error

	if err != nil {
		return Book{}, err
	}

	return book, err
}

func (r *repository) CreateBook(book Book) (Book, error) {
	err := r.db.Create(&book).Error

	if err != nil {
		return Book{}, err
	}

	return book, err
}

func (r *repository) DeleteBook(id int) (Book, error) {
	var book Book
	err := r.db.First(&book, id).Error

	if err != nil {
		return Book{}, err
	}

	err = r.db.Delete(&book).Error

	if err != nil {
		return Book{}, err
	}

	return book, err
}

func (r *repository) UpdateBook(book Book) (Book, error) {
	err := r.db.Save(&book).Error

	if err != nil {
		return Book{}, err
	}

	return book, err
}
