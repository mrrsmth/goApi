package services

import (
	"go_api_template/pkg/entities"
	"go_api_template/pkg/repositories"
)

type IUserService interface {
	GetUserByID(id int) (*entities.User, error)
}

// UserService contains the business logic for users.
type UserService struct {
	usersRepo repositories.IUserRepo
}

// NewUserService creates a new UserService.
func NewUserService(usersRepo repositories.IUserRepo) *UserService {
	return &UserService{usersRepo: usersRepo}
}

// GetUserByID returns a user by their ID.
func (s *UserService) GetUserByID(id int) (*entities.User, error) {
	return s.usersRepo.GetByID(id)
}
