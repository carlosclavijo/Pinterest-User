package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/commands"
	"github.com/google/uuid"
)

func (h *UserHandler) HandleUpdateProfilePic(ctx context.Context, cmd commands.UpdateProfilePicCommand) error {
	id, err := uuid.Parse(cmd.UserID)
	if err != nil {
		return err
	}

	user, err := h.repository.GetById(ctx, id)
	if err != nil {
		return err
	}

	user.ChangeProfilePic(&cmd.ProfilePic)
	_, err = h.repository.Update(ctx, user)
	return err
}
