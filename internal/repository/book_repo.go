package repository

import (
	"LibaryBookControl/internal/models"
	"database/sql"
	"errors"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) CreateBook(book *models.Book) error {

	query := "INSERT INTO books (title, year, author_id) VALUES ($1, $2, $3) RETURNING id"

	err := r.db.QueryRow(query, book.Title, book.Year, book.AuthorID).Scan(&book.ID)

	return err

}

func (r *BookRepository) GetAllBooks(book *models.Book) ([]models.Book, error) {

	query := "SELECT id, title, year, author_id FROM books"

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Year, &book.AuthorID)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil

}

func (r *BookRepository) GetBookByID(id int) (*models.Book, error) {

	query := "SELECT title, year, author_id FROM books WHERE id = $1"

	row := r.db.QueryRow(query, id)

	var book models.Book

	err := row.Scan(&book.ID, &book.Title, &book.Year, &book.AuthorID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &book, err

}

func (r *BookRepository) UpdateBook(book *models.Book) error {

	query := "UPDATE books SET title = $1, year = $2, author_id = $3 WHERE id = $4"

	result, err := r.db.Exec(query, book.Title, book.Year, book.AuthorID, book.ID)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Book not found")
	}

	return err

}

func (r *BookRepository) DeleteBook(book *models.Book) error {

	query := "DELETE FROM books WHERE id = $1"

	result, err := r.db.Exec(query, book.ID)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Book not found")
	}

	return err

}
