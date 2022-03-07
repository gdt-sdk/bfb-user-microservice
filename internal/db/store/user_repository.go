package store

import (
	"database/sql"
)

type UserRepository struct {
	store *sql.DB
}

// func (r *UserRepository) Create(u *models.User) error {
// 	if err := r.store.QueryRow(
// 		`
// 		INSERT INTO users()
// 		`,
// 	)
// }
