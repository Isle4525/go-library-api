package repository

import (
	"LibaryBookControl/internal/models"
	"database/sql"
)

type BorrowRepository struct {
	db *sql.DB
}

func NewBorrowRepository(db *sql.DB) *BorrowRepository {
	return &BorrowRepository{db: db}
}

func (r *BorrowRepository) CreateBorrowRecord(record *models.BorrowRecord) error {
	query := "INSERT INTO borrow_records (book_id, user_id, borrow_date, return_date, is_returned) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	err := r.db.QueryRow(query, record.BookID, record.UserID, record.BorrowDate, record.ReturnDate, record.IsReturned).Scan(&record.ID)

	return err
}

func (r *BorrowRepository) GetAllBorrowRecords() ([]models.BorrowRecord, error) {
	query := "SELECT id, book_id, user_id, borrow_date, return_date, is_returned FROM borrow_records"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []models.BorrowRecord
	for rows.Next() {
		var record models.BorrowRecord
		err := rows.Scan(&record.ID, &record.BookID, &record.UserID, &record.BorrowDate, &record.ReturnDate, &record.IsReturned)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (r *BorrowRepository) GetBorrowRecordByID(id int) (*models.BorrowRecord, error) {
	query := "SELECT id, book_id, user_id, borrow_date, return_date, is_returned FROM borrow_records WHERE id = $1"

	row := r.db.QueryRow(query, id)

	var record models.BorrowRecord
	err := row.Scan(&record.ID, &record.BookID, &record.UserID, &record.BorrowDate, &record.ReturnDate, &record.IsReturned)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &record, nil

}
