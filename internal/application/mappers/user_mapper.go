package mappers

import (
	"github.com/carlosclavijo/Pinterest-User/internal/application/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/user"
	"time"
)

func MapToUserDTO(user *users.User) *dto.UserDTO {
	var phone, information, profilePic, website *string

	if user.Phone() != nil {
		p := user.Phone().String()
		phone = &p
	}

	if user.Information() != nil {
		information = user.Information()
	}

	if user.ProfilePic() != nil {
		profilePic = user.ProfilePic()
	}

	if user.WebSite() != nil {
		w := user.WebSite().String()
		website = &w
	}

	return &dto.UserDTO{
		Id:          user.Id(),
		FirstName:   user.FirstName(),
		LastName:    user.LastName(),
		Username:    user.Username().String(),
		Email:       user.Email().String(),
		Gender:      user.Gender().String(),
		Birth:       user.Birth().Time(),
		Country:     user.Country().String(),
		Language:    user.Language().String(),
		Phone:       phone,
		Information: information,
		ProfilePic:  profilePic,
		Website:     website,
		Visibility:  user.Visibility(),
	}
}

func MapToUserResponse(user *dto.UserDTO, lastLoginAt, createdAt, updatedAt time.Time, deletedAt *time.Time) *dto.UserResponse {
	return &dto.UserResponse{
		UserDTO:     user,
		LastLoginAt: lastLoginAt,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   deletedAt,
	}
}
