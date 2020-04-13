package store

import "../model"

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u model.User) (model.User, error) {
	return nil, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) {
	return nil, nil
}
