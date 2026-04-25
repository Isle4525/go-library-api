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

	author, err := s.repo.GetAuthorByID(id)

	if err != nil {
		return nil, err
	}

	if author == nil {
		return nil, errors.New("Author not found")
	}

	return author, nil
}

func (s *AuthorService) UpdateAuthor(author *models.Author) error {

	if author.ID <= 0 {
		return errors.New("Uncorrect ID")
	}

	author, err := s.repo.GetAuthorByID(author.ID)
	if err != nil {
		return err
	}

	if author == nil {
		return errors.New("Author not found")
	}

	return s.repo.UpdateAuthor(author)
}

func (s *AuthorService) DeleteAuthor(id int) error {

	if id <= 0 {
		return errors.New("Uncorrect ID")
	}

	author, err := s.repo.GetAuthorByID(id)
	if err != nil {
		return err
	}

	if author == nil {
		return errors.New("Author not found")
	}

	return s.repo.DeleteAuthor(id)
}
