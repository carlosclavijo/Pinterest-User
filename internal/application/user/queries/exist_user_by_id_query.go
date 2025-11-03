package queries

import "github.com/google/uuid"

type ExistUserByIdQuery struct {
	Id uuid.UUID `json:"id"`
}
