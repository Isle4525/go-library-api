package models

type BorrowRecord struct {
	ID         int    `json:"id"`
	BookID     int    `json:"book_id"`
	UserID     int    `json:"user_id"`
	BorrowDate string `json:"borrow_date"`
	ReturnDate string `json:"return_date"`
	IsReturned bool   `json:"is_returned"`
}
