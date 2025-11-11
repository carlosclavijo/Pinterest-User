package repositories

import (
	"context"
	"database/sql"
	"time"
)

type emailVerificationRepo struct {
	db *sql.DB
}

func (e *emailVerificationRepo) Save(ctx context.Context, ev *email.EmailVerification) error {
	query := `
        INSERT INTO email_verifications (id, user_id, token, created_at, expires_at)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := e.db.ExecContext(ctx, query, ev.Id, ev.UserId, ev.Token, ev.CreatedAt, ev.ExpiresAt)
	return err
}

func (e *emailVerificationRepo) FindByToken(ctx context.Context, token string) (*email.EmailVerification, error) {
	query := `
        SELECT id, user_id, token, created_at, expires_at, verified_at
        FROM email_verifications
        WHERE token = $1
    `
	ev := &email.EmailVerification{}
	row := e.db.QueryRowContext(ctx, query, token)
	err := row.Scan(&ev.Id, &ev.UserId, &ev.Token, &ev.CreatedAt, &ev.ExpiresAt, &ev.VerifiedAt)
	if err != nil {
		return nil, err
	}
	return ev, nil
}

func (e *emailVerificationRepo) MarkVerified(ctx context.Context, ev *email.EmailVerification) error {
	query := `
        UPDATE email_verifications
        SET verified_at = $1
        WHERE id = $2
    `
	_, err := e.db.ExecContext(ctx, query, time.Now(), ev.Id)
	return err
}

func NewEmailVerificationRepo(db *sql.DB) email.EmailVerificationRepository {
	return &emailVerificationRepo{db: db}
}
