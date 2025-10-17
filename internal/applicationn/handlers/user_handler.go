package handlers

import "github.com/carlosclavijo/Pinterest-User/internal/domain/user"

type UserHandler struct {
	repository user.UserRepository
	factory    user.UserFactory
}

func NewUserHandler(repository user.UserRepository, factory user.UserFactory) *UserHandler {
	return &UserHandler{
		repository: repository,
		factory:    factory,
	}
}
