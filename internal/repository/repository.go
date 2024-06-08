package repository

import (
	"github.com/nurhidaylma/lover-app.git/internal/repository/interfaces"
)

type Repository struct {
	DB    interfaces.DBRepository
	Redis interfaces.RedisRepository
}

type InitRepo interface {
	NewRepository(*Repository) error
}

func NewLoverRepository(repos []InitRepo) (*Repository, error) {
	repo := &Repository{}

	for _, r := range repos {
		if err := r.NewRepository(repo); err != nil {
			return nil, err
		}
	}

	return repo, nil
}
