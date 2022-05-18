package product

import "gorm.io/gorm"

type Repository interface {
	GetProducts() ([]Product, error)
	GetProduct(id int) (Product, error)
	CreateProduct(product Product) (Product, error)
	DeleteProduct(id int) (Product, error)
	UpdateProduct(product Product) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProducts() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, err
}

func (r *repository) GetProduct(id int) (Product, error) {
	var product Product
	err := r.db.First(&product, id).Error

	if err != nil {
		return Product{}, err
	}

	return product, err
}

func (r *repository) CreateProduct(product Product) (Product, error) {
	err := r.db.Create(&product).Error

	if err != nil {
		return Product{}, err
	}

	return product, err
}

func (r *repository) DeleteProduct(id int) (Product, error) {
	var product Product
	err := r.db.First(&product, id).Error

	if err != nil {
		return Product{}, err
	}

	err = r.db.Delete(&product).Error

	if err != nil {
		return Product{}, err
	}

	return product, err
}

func (r *repository) UpdateProduct(product Product) (Product, error) {
	err := r.db.Save(&product).Error

	if err != nil {
		return Product{}, err
	}

	return product, err
}
