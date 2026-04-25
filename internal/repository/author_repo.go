package repository

import (
	"LibaryBookControl/internal/models"
	"database/sql"
	"errors"
)

type AuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) CreateAuthor(author *models.Author) error {
	query := "INSERT INTO authors (name, bio) VALUES ($1, $2) RETURNING id"

	err := r.db.QueryRow(query, author.Name, author.Bio).Scan(&author.ID)

	return err
}

func (r *AuthorRepository) GetAllAuthors() ([]models.Author, error) {
	query := "SELECT id, name, bio FROM authors"

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []models.Author

	for rows.Next() {
		var author models.Author
		err := rows.Scan(&author.ID, &author.Name, &author.Bio)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, nil

}

func (r *AuthorRepository) GetAuthorByID(id int) (*models.Author, error) {

	query := "SELECT name, bio FROM authors WHERE id = $1"

	row := r.db.QueryRow(query, id)

	var author models.Author

	err := row.Scan(&author.ID, &author.Name, &author.Bio)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &author, nil

}

func (r *AuthorRepository) UpdateAuthor(author *models.Author) error {
	query := "UPDATE authors SET name = $1, bio = $2 WHERE id = $3"

	result, err := r.db.Exec(query, author.Name, author.Bio, author.ID)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Author not found")
	}

	return err

}

func (r *AuthorRepository) DeleteAuthor(id int) error {
	query := "DELETE FROM authors WHERE id = $1"

	result, err := r.db.Exec(query, id)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Author not found")
	}

	return err

}
