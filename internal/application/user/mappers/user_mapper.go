package mappers

import (
	"fmt"
	"github.com/carlosclavijo/Pinterest-Services/internal/application/user/dto"
	users "github.com/carlosclavijo/Pinterest-Services/internal/domain/user"
	"time"
)

func MapToUserDTO(user *users.User) *dto.UserDTO {
	var phone, information, profilePic, profilePicURL, website *string

	if user.Phone() != nil {
		p := user.Phone().String()
		phone = &p
	}

	if user.Information() != nil {
		information = user.Information()
	}

	if user.ProfilePic() != nil {
		profilePic = user.ProfilePic()
		url := fmt.Sprintf("http://localhost:8080/static/%s", *user.ProfilePic())
		profilePicURL = &url
	}

	if user.WebSite() != nil {
		w := user.WebSite().String()
		website = &w
	}

	return &dto.UserDTO{
		Id:            user.Id(),
		FirstName:     user.FirstName(),
		LastName:      user.LastName(),
		Username:      user.Username().String(),
		Email:         user.Email().String(),
		Gender:        user.Gender().String(),
		Birth:         user.Birth().Time(),
		Country:       user.Country().String(),
		Language:      user.Language().String(),
		Phone:         phone,
		Information:   information,
		ProfilePic:    profilePic,
		ProfilePicURL: profilePicURL,
		Website:       website,
		Visibility:    user.Visibility(),
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
