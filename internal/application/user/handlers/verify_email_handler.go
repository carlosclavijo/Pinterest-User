package handlers

import (
	"context"
	"errors"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/commands"
)

func (h *UserHandler) HandleVerifyEmail(ctx context.Context, cmd commands.VerifyEmailCommand) error {
	ev, err := h.emailRepo.FindByToken(ctx, cmd.Token)
	if err != nil {
		return err
	}

	if ev.IsExpired() {
		return errors.New("token expired")
	}

	if ev.IsVerified() {
		return errors.New("already verified")
	}

	ev.MarkVerified()
	return h.emailRepo.Save(ctx, ev)
}
