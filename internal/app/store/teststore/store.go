package teststore

import (
	"database/sql"

	store ".."
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
