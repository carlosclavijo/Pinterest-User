package queries

import "github.com/google/uuid"

type GetListByUserIdQuery struct {
	UserId uuid.UUID `json:"user_id"`
}
