package email

import (
	"github.com/google/uuid"
	"time"
)

type EmailVerification struct {
	Id         uuid.UUID  `json:"id"`
	UserId     uuid.UUID  `json:"user_id"`
	Token      string     `json:"token"`
	CreatedAt  time.Time  `json:"created_at"`
	ExpiresAt  time.Time  `json:"expires_at"`
	VerifiedAt *time.Time `json:"verified_at,omitempty"`
}

func (ev *EmailVerification) IsExpired() bool {
	return time.Now().After(ev.ExpiresAt)
}

func (ev *EmailVerification) IsVerified() bool {
	return ev.VerifiedAt != nil
}

func (ev *EmailVerification) MarkVerified() {
	now := time.Now()
	ev.VerifiedAt = &now
}
