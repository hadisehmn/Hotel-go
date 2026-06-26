package repository

import (
	"database/sql"
	"fmt"
	"go-practice/HOTEL/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) CreateUser(u models.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users(name, phone, password_hash) VALUES ($1, $2, $3)",
		u.Name,
		u.Phone,
		u.Password,
	)
	return err
}

func (r *UserRepository) FindByName(name string) (models.User, error) {
	var user models.User
	err := r.DB.QueryRow(
		"SELECT name, password_hash FROM users WHERE name = $1", name,
	).Scan(
		&user.Name,
		&user.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, err
	}

	return user, nil

}

func (r *UserRepository) ExistsByName(name string) (bool, error) {
	var exists bool

	err := r.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE name=$1)",
		name,
	).Scan(&exists)

	return exists, err
}
