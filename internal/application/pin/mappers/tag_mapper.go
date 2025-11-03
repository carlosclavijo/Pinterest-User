package mappers

import (
	"github.com/carlosclavijo/Pinterest-User/internal/application/pin/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/tag"
	"time"
)

func MapToTagDTO(tag *tag.Tag) *dto.TagDTO {
	return &dto.TagDTO{
		Id:   tag.Id(),
		Name: tag.Name(),
	}
}

func MapToTagResponse(tag *dto.TagDTO, createdAt time.Time, deletedAt *time.Time) *dto.TagResponse {
	return &dto.TagResponse{
		TagDTO:    tag,
		CreatedAt: createdAt,
		DeletedAt: deletedAt,
	}
}
