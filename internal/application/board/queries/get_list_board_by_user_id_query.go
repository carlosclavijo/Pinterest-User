package queries

import "github.com/google/uuid"

type GetListBoardsByUserIdQuery struct {
	UserId uuid.UUID `json:"user_id"`
}
