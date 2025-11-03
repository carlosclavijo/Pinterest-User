package queries

import "github.com/google/uuid"

type GetPinByIdQuery struct {
	Id uuid.UUID `json:"id"`
}
