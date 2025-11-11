package mappers

import (
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/shared"
	users "github.com/carlosclavijo/Pinterest-Services/internal/domain/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestMapToUserDTO(t *testing.T) {
	firstName := "John"
	lastName := "Doe"
	userNameStr := "johndoes"
	emailStr := "john@doe.com"
	passwordStr := "5Trong!."
	genderStr := "Male"
	birthTm := time.Now().AddDate(-20, 0, 1)
	countryStr := "Bolivia"
	languageStr := "Spanish"
	phoneStr := "+591-70926048"
	informationStr := "info"
	profilePicStr := "./images/id.jpg"
	webSiteStr := "https://www.github.com/carlosclavijo/"

	userName, err := shared.NewUsername(userNameStr)
	assert.NotEmpty(t, userName)
	assert.NoError(t, err)

	email, err := shared.NewEmail(emailStr)
	assert.NotEmpty(t, email)
	assert.NoError(t, err)

	password, err := shared.NewPassword(passwordStr)
	assert.NotEmpty(t, password)
	assert.NoError(t, err)

	gender, err := shared.ParseGender(genderStr)
	assert.NotEmpty(t, gender)
	assert.NoError(t, err)

	birth, err := shared.NewBirthDate(birthTm)
	assert.NotEmpty(t, birth)
	assert.NoError(t, err)

	country, err := shared.ParseCountry(countryStr)
	assert.NotEmpty(t, country)
	assert.NoError(t, err)

	language, err := shared.ParseLanguage(languageStr)
	assert.NotEmpty(t, language)
	assert.NoError(t, err)

	usr := users.NewUser(firstName, lastName, userName, email, password, gender, birth, country, language, nil)
	require.NotEmpty(t, usr)

	userDTO := MapToUserDTO(usr)

	require.NotEmpty(t, userDTO)
	assert.Equal(t, usr.Id(), userDTO.Id)
	assert.Equal(t, usr.FirstName(), userDTO.FirstName)
	assert.Equal(t, usr.LastName(), userDTO.LastName)
	assert.Equal(t, usr.Username().String(), userDTO.Username)
	assert.Equal(t, usr.Email().String(), userDTO.Email)
	assert.Equal(t, usr.Gender().String(), userDTO.Gender)
	assert.Equal(t, usr.Birth().Time(), userDTO.Birth)
	assert.Equal(t, usr.Country().String(), userDTO.Country)
	assert.Equal(t, usr.Language().String(), userDTO.Language)
	assert.Nil(t, userDTO.Phone)
	assert.Nil(t, userDTO.Information)
	assert.Nil(t, userDTO.ProfilePic)
	assert.Nil(t, userDTO.Website)
	assert.False(t, userDTO.Visibility)

	phone, err := shared.NewPhone(&phoneStr)
	assert.NotEmpty(t, phone)
	assert.NoError(t, err)

	webSite, err := shared.NewWebSite(&webSiteStr)
	assert.NotEmpty(t, webSite)
	assert.NoError(t, err)

	err = usr.ChangeInformation(&informationStr)

	assert.NoError(t, err)

	usr.ChangePhone(phone)
	usr.ChangeProfilePic(&profilePicStr)
	usr.ChangeWebSite(webSite)
	usr.ChangeVisibility(true)

	userDTO = MapToUserDTO(usr)

	assert.NotNil(t, usr.Phone())
	assert.Equal(t, usr.Phone().String(), *userDTO.Phone)
	assert.NotNil(t, usr.Information())
	assert.Equal(t, informationStr, *userDTO.Information)
	assert.NotNil(t, usr.ProfilePic())
	assert.Equal(t, profilePicStr, *userDTO.ProfilePic)
	assert.NotNil(t, usr.WebSite())
	assert.Equal(t, webSiteStr, *userDTO.Website)
	assert.True(t, usr.Visibility())

	lastLoginAt := time.Now()
	createdAt := time.Now()
	updatedAt := time.Now()

	userResponse := MapToUserResponse(userDTO, lastLoginAt, createdAt, updatedAt, nil)

	require.NotEmpty(t, userResponse)
	assert.Exactly(t, userDTO, userResponse.UserDTO)
	assert.Equal(t, lastLoginAt, userResponse.LastLoginAt)
	assert.WithinDuration(t, time.Now(), userResponse.LastLoginAt, time.Second)
	assert.Equal(t, createdAt, userResponse.CreatedAt)
	assert.WithinDuration(t, time.Now(), userResponse.CreatedAt, time.Second)
	assert.Equal(t, updatedAt, userResponse.UpdatedAt)
	assert.WithinDuration(t, time.Now(), userResponse.UpdatedAt, time.Second)
	assert.Nil(t, userResponse.DeletedAt)
}
