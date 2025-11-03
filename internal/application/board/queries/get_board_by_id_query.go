package queries

import "github.com/google/uuid"

type GetBoardByIdQuery struct {
	Id uuid.UUID `json:"id"`
}
