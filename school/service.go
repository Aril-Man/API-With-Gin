package school

type Service interface {
	GetSchools() ([]School, error)
	GetSchool(id int) (School, error)
	CreateSchool(schoolRequest SchoolRequest) (School, error)
	DeleteSchool(id int) (School, error)
	UpdateSchool(ID int, schoolRequest SchoolRequest) (School, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetSchools() ([]School, error) {
	school, err := s.repo.GetSchools()

	if err != nil {
		return nil, err
	}

	return school, nil
}

func (s *service) GetSchool(id int) (School, error) {
	school, err := s.repo.GetSchool(id)

	if err != nil {
		return School{}, err
	}

	return school, nil
}

func (s *service) CreateSchool(schoolRequest SchoolRequest) (School, error) {

	school := School{
		Name:    schoolRequest.Name,
		Address: schoolRequest.Address,
		Class:   schoolRequest.Class,
		Major:   schoolRequest.Major,
	}

	school, err := s.repo.CreateSchool(school)

	if err != nil {
		return School{}, err
	}

	return school, nil
}

func (s *service) DeleteSchool(id int) (School, error) {
	school, err := s.repo.DeleteSchool(id)

	if err != nil {
		return School{}, err
	}

	return school, nil
}

func (s *service) UpdateSchool(ID int, schoolRequest SchoolRequest) (School, error) {

	school, err := s.repo.GetSchool(ID)

	if err != nil {
		return School{}, err
	}

	school.Name = schoolRequest.Name
	school.Address = schoolRequest.Address
	school.Class = schoolRequest.Class
	school.Major = schoolRequest.Major

	NewSchool, err := s.repo.UpdateSchool(school)

	if err != nil {
		return School{}, err
	}

	return NewSchool, nil
}
