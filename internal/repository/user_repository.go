package repository

import (
	"github.com/jmoiron/sqlx"
	"golang-marketplace/internal/entity"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(request *entity.User) error {
	query := `INSERT INTO users ( name, username, password) VALUES ($1, $2, $3)`

	_, err := r.DB.Exec(query, request.Name, request.Username, request.Password)
	return err
}
