package queries

import "github.com/google/uuid"

type GetUserByIdQuery struct {
	Id uuid.UUID `json:"id"`
}
