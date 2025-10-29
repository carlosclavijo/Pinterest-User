package queries

import "github.com/google/uuid"

type GetUserById struct {
	Id uuid.UUID `json:"id"`
}
