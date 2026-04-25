package repository

import (
	"LibaryBookControl/internal/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {

	query := "INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id"

	err := r.db.QueryRow(query, user.Username, user.Email).Scan(&user.ID)

	return err

}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {

	query := "SELECT id, username, email FROM users"

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	query := "SELECT username, email FROM users WHERE id = $1"

	row := r.db.QueryRow(query, id)

	var user models.User

	err := row.Scan(&user.ID, &user.Username, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil

}

func (r *UserRepository) UpdateUser(user *models.User) error {
	query := "UPDATE users SET username = $1, email = $2 WHERE id = $3"

	result, err := r.db.Exec(query, user.Username, user.Email, user.ID)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return err
	}

	return err

}

func (r *UserRepository) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = $1"

	result, err := r.db.Exec(query, id)

	resultAffeted, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if resultAffeted == 0 {
		return err
	}

	return err

}
