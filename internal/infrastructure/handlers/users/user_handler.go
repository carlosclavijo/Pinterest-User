package users

import users "github.com/carlosclavijo/Pinterest-Services/internal/domain/user"

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
