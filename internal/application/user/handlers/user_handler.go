package handlers

import (
	"github.com/carlosclavijo/Pinterest-Services/internal/application"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/email"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/user"
)

type UserHandler struct {
	repository   users.UserRepository
	emailRepo    email.EmailVerificationRepository
	emailService email.Sender
	factory      users.UserFactory
	logger       application.Logger
}

func NewUserHandler(repository users.UserRepository, emailRepo email.EmailVerificationRepository, emailService email.Sender, factory users.UserFactory, logger application.Logger) *UserHandler {
	return &UserHandler{
		repository:   repository,
		emailRepo:    emailRepo,
		emailService: emailService,
		factory:      factory,
		logger:       logger,
	}
}
