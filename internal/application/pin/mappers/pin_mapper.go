package mappers

import (
	"github.com/carlosclavijo/Pinterest-User/internal/application/pin/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/pin"
	"time"
)

func MapToPinDTO(pin *pins.Pin) *dto.PinDTO {
	var tagsDTO []*dto.TagDTO
	for _, k := range pin.Tags() {
		t := MapToTagDTO(&k)
		tagsDTO = append(tagsDTO, t)
	}
	return &dto.PinDTO{
		Id:           pin.Id(),
		UserId:       pin.UserId(),
		BoardId:      pin.BoardId(),
		Title:        pin.Title(),
		Description:  pin.Description(),
		Image:        pin.Image(),
		SaveCount:    pin.SaveCount(),
		LikeCount:    pin.LikeCount(),
		CommentCount: pin.CommentCount(),
		Visibility:   pin.Visibility(),
		Tags:         tagsDTO,
	}
}

func MapToPinResponse(pin *dto.PinDTO, createdAt, updatedAt time.Time, deletedAt *time.Time) *dto.PinResponse {
	return &dto.PinResponse{
		PinDTO:    pin,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
