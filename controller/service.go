package controller

import "technical-test-02-10-22/models"

type Service interface {
	Create(input FileInput) (models.File, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r: r}
}

func (s *service) Create(input FileInput) (models.File, error) {
	fileCreate := models.File{}

	fileCreate.FileName = input.NameFile
	fileCreate.Size = input.Size
	fileCreate.ContentType = input.ContentType

	newFile, err := s.r.Create(fileCreate)
	if err != nil {
		return newFile, err
	}

	return newFile, nil
}
