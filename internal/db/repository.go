package db

import "github.com/gefion-tech/bfb-user-microservice/internal/models"

type UserRepository interface {
	Create(u *models.User) error
}
