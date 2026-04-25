package services

import (
	"LibaryBookControl/internal/models"
	"LibaryBookControl/internal/repository"
	"errors"
)

type AuthorService struct {
	repo *repository.AuthorRepository
}

func NewAuthorService(repo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) CreateAuthor(author *models.Author) error {
	return s.repo.CreateAuthor(author)
}

func (s *AuthorService) GetAllAuthors() ([]models.Author, error) {
	return s.repo.GetAllAuthors()
}

func (s *AuthorService) GetAuthorByID(id int) (*models.Author, error) {

	if id <= 0 {
		return nil, errors.New("Uncorrect ID")
	}

	book, err := s.repo.GetAuthorByID(id)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *AuthorService) UpdateAuthor(author *models.Author) error {
	return s.repo.UpdateAuthor(author)
}

func (s *AuthorService) DeleteAuthor(id int) error {
	return s.repo.DeleteAuthor(id)
}
