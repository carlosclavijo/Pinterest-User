package email

import "context"

type EmailVerificationRepository interface {
	Save(ctx context.Context, ev *EmailVerification) error
	FindByToken(ctx context.Context, token string) (*EmailVerification, error)
	MarkVerified(ctx context.Context, ev *EmailVerification) error
}
