package product

type Service interface {
	GetProducts() ([]Product, error)
	GetProduct(id int) (Product, error)
	CreateProduct(productRequest ProductRequest) (Product, error)
	DeleteProduct(id int) (Product, error)
	UpdateProduct(ID int, productRequest ProductRequest) (Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetProducts() ([]Product, error) {
	product, err := s.repo.GetProducts()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *service) GetProduct(id int) (Product, error) {
	product, err := s.repo.GetProduct(id)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) CreateProduct(productRequest ProductRequest) (Product, error) {

	product := Product{
		Product: productRequest.Product,
		Price:   productRequest.Price,
	}

	product, err := s.repo.CreateProduct(product)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) DeleteProduct(id int) (Product, error) {
	product, err := s.repo.DeleteProduct(id)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) UpdateProduct(ID int, productrequest ProductRequest) (Product, error) {

	product, err := s.repo.GetProduct(ID)

	if err != nil {
		return Product{}, err
	}

	product.Product = productrequest.Product
	product.Price = productrequest.Price

	NewProduct, err := s.repo.UpdateProduct(product)

	if err != nil {
		return Product{}, err
	}

	return NewProduct, nil
}
