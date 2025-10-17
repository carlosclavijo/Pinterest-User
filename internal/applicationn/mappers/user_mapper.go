package mappers

import (
	"github.com/carlosclavijo/Pinterest-User/internal/applicationn/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/user"
	"time"
)

func MapToUserDTO(user *user.User) *dto.UserDTO {
	var phone *string
	if user.Phone() != nil {
		p := user.Phone().String()
		phone = p
	}

	return &dto.UserDTO{
		Id:        user.Id(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Username:  user.Username().String(),
		Email:     user.Email().String(),
		Gender:    user.Gender().String(),
		Birth:     user.Birth().Time(),
		Country:   user.Country().String(),
		Language:  user.Language().String(),
		Phone:     phone,
	}
}

func MapToUserResponse(user *dto.UserDTO, createdAt, updatedAt time.Time, deletedAt *time.Time) *dto.UserResponse {
	return &dto.UserResponse{
		UserDTO:   user,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
