package store

import (
	"database/sql"

	"github.com/gefion-tech/bfb-user-microservice/internal/db"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func InitStore(db *sql.DB) db.SQLStoreI {
	return &Store{
		db: db,
	}
}

func (s *Store) User() db.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s.db,
	}

	return s.userRepository
}
