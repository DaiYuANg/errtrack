package repository

import "github.com/uptrace/bun"

type UserRepository struct {
	*bun.DB
}

func newUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db}
}
