package book

type Service interface {
	GetBooks() ([]Book, error)
	GetBook(id int) (Book, error)
	CreateBook(bookRequest BookRequest) (Book, error)
	DeleteBook(id int) (Book, error)
	UpdateBook(ID int, bookRequest BookRequest) (Book, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetBooks() ([]Book, error) {
	books, err := s.repo.GetBooks()

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *service) GetBook(id int) (Book, error) {
	book, err := s.repo.GetBook(id)

	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func (s *service) CreateBook(bookRequest BookRequest) (Book, error) {

	book := Book{
		Title:       bookRequest.Title,
		Price:       bookRequest.Price,
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
	}

	book, err := s.repo.CreateBook(book)

	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func (s *service) DeleteBook(id int) (Book, error) {
	book, err := s.repo.DeleteBook(id)

	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func (s *service) UpdateBook(ID int, bookRequest BookRequest) (Book, error) {

	book, err := s.repo.GetBook(ID)

	if err != nil {
		return Book{}, err
	}

	book.Title = bookRequest.Title
	book.Price = bookRequest.Price
	book.Description = bookRequest.Description
	book.Rating = bookRequest.Rating

	NewBook, err := s.repo.UpdateBook(book)

	if err != nil {
		return Book{}, err
	}

	return NewBook, nil
}
