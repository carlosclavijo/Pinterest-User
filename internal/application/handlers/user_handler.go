package handlers

import "github.com/carlosclavijo/Pinterest-User/internal/domain/user"

type UserHandler struct {
	repository users.UserRepository
	factory    users.UserFactory
}

func NewUserHandler(repository users.UserRepository, factory users.UserFactory) *UserHandler {
	return &UserHandler{
		repository: repository,
		factory:    factory,
	}
}
