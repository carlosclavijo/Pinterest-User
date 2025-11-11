package handlers

import (
	"context"
	"fmt"
	"time"
)

func (h *UserHandler) VerifyEmailToken(ctx context.Context, token string) error {
	ev, err := h.emailRepo.FindByToken(ctx, token)
	if err != nil || ev == nil {
		return fmt.Errorf("invalid or expired token")
	}

	// Mark it as verified
	now := time.Now()
	ev.VerifiedAt = &now

	if err := h.emailRepo.MarkVerified(ctx, ev); err != nil {
		return fmt.Errorf("failed to mark token verified: %w", err)
	}

	return nil
}
