package mappers

import (
	"github.com/carlosclavijo/Pinterest-Services/internal/application/pin/dto"
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/pin"
	"time"
)

func MapToTagDTO(tag *pins.Tag) *dto.TagDTO {
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
