package school

import "gorm.io/gorm"

type Repository interface {
	GetSchools() ([]School, error)
	GetSchool(id int) (School, error)
	CreateSchool(school School) (School, error)
	DeleteSchool(id int) (School, error)
	UpdateSchool(school School) (School, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetSchools() ([]School, error) {
	var schools []School
	err := r.db.Find(&schools).Error

	if err != nil {
		return nil, err
	}

	return schools, err
}

func (r *repository) GetSchool(id int) (School, error) {
	var school School
	err := r.db.First(&school, id).Error

	if err != nil {
		return School{}, err
	}

	return school, err
}

func (r *repository) CreateSchool(school School) (School, error) {
	err := r.db.Create(&school).Error

	if err != nil {
		return School{}, err
	}

	return school, err
}

func (r *repository) DeleteSchool(id int) (School, error) {
	var school School
	err := r.db.First(&school, id).Error

	if err != nil {
		return School{}, err
	}

	err = r.db.Delete(&school).Error

	if err != nil {
		return School{}, err
	}

	return school, err
}

func (r *repository) UpdateSchool(school School) (School, error) {
	err := r.db.Save(&school).Error

	if err != nil {
		return School{}, err
	}

	return school, err
}
