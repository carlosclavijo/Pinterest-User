package users

import (
	"github.com/carlosclavijo/Pinterest-User/internal/domain/shared"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
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

	phone, err := shared.NewPhone(&phoneStr)
	assert.NotEmpty(t, phone)
	assert.NoError(t, err)

	usr := NewUser(firstName, lastName, userName, email, password, gender, birth, country, language, phone)

	require.NotEmpty(t, usr)
	assert.NotNil(t, usr.Id())
	assert.Equal(t, firstName, usr.FirstName())
	assert.Equal(t, lastName, usr.LastName())
	assert.Exactly(t, userName, usr.Username())
	assert.Equal(t, userNameStr, usr.Username().String())
	assert.Exactly(t, email, usr.Email())
	assert.Equal(t, emailStr, usr.Email().String())
	assert.Exactly(t, password, usr.Password())
	assert.Equal(t, passwordStr, usr.Password().String())
	assert.Exactly(t, gender, usr.Gender())
	assert.Equal(t, genderStr, usr.Gender().String())
	assert.Exactly(t, birth, usr.Birth())
	assert.Equal(t, birthTm.Format(time.RFC3339), usr.Birth().Time().Format(time.RFC3339))
	assert.Exactly(t, country, usr.Country())
	assert.Equal(t, countryStr, usr.Country().String())
	assert.Exactly(t, language, usr.Language())
	assert.Equal(t, languageStr, usr.Language().String())
	assert.Exactly(t, phone, usr.Phone())
	assert.Equal(t, phoneStr, usr.Phone().String())
	assert.Nil(t, usr.Information())
	assert.Nil(t, usr.ProfilePic())
	assert.Nil(t, usr.WebSite())
	assert.False(t, usr.Visibility())
	assert.WithinDuration(t, time.Now(), usr.LastLoginAt(), time.Second)
	assert.WithinDuration(t, time.Now(), usr.CreatedAt(), time.Second)
	assert.WithinDuration(t, time.Now(), usr.UpdatedAt(), time.Second)
	assert.Nil(t, usr.DeletedAt())

	usr = NewUser(firstName, lastName, userName, email, password, gender, birth, country, language, nil)

	assert.Nil(t, usr.Phone())

	firstName = ""
	err = usr.ChangeFirstName(firstName)

	assert.ErrorIs(t, err, ErrEmptyFirstNameUser)

	firstName = "averylongfirstnamewhichofcoursewouldreturnanerrorandhavemorethanahundredcharacterslongtohavethaterror"
	err = usr.ChangeFirstName(firstName)

	assert.ErrorIs(t, err, ErrLongFirstNameUser)

	firstName = "John Doe"
	err = usr.ChangeFirstName(firstName)

	assert.Equal(t, firstName, usr.FirstName())
	assert.NoError(t, err)

	lastName = ""
	err = usr.ChangeLastName(lastName)

	assert.ErrorIs(t, err, ErrEmptyLastNameUser)

	lastName = "averylongfirstnamewhichofcoursewouldreturnanerrorandhavemorethanahundredcharacterslongtohavethaterror"
	err = usr.ChangeLastName(lastName)

	assert.ErrorIs(t, err, ErrLongLastNameUser)

	lastName = "John Doe"
	err = usr.ChangeLastName(lastName)

	assert.Equal(t, lastName, usr.LastName())
	assert.NoError(t, err)

	newUserName := shared.Username{}
	err = usr.ChangeUsername(newUserName)

	assert.ErrorIs(t, err, shared.ErrEmptyUsername)

	err = usr.ChangeUsername(userName)

	assert.Exactly(t, userName, usr.Username())
	assert.Equal(t, userNameStr, usr.Username().String())
	assert.NoError(t, err)

	newEmail := shared.Email{}
	err = usr.ChangeEmail(newEmail)

	assert.ErrorIs(t, err, shared.ErrEmptyEmail)

	err = usr.ChangeEmail(email)

	assert.Exactly(t, email, usr.Email())
	assert.Equal(t, emailStr, usr.Email().String())
	assert.NoError(t, err)

	newPassword := shared.Password{}
	err = usr.ChangePassword(newPassword)

	assert.ErrorIs(t, err, shared.ErrEmptyPassword)

	err = usr.ChangePassword(password)

	assert.Exactly(t, password, usr.Password())
	assert.Equal(t, passwordStr, usr.Password().String())
	assert.NoError(t, err)

	newGender := shared.Gender("X")
	err = usr.ChangeGender(newGender)

	assert.ErrorIs(t, err, shared.ErrNotAGender)

	err = usr.ChangeGender(gender)

	assert.Exactly(t, gender, usr.Gender())
	assert.Equal(t, genderStr, usr.Gender().String())
	assert.NoError(t, err)

	newBirth := shared.BirthDate{}
	err = usr.ChangeBirth(newBirth)

	assert.ErrorIs(t, err, shared.ErrEmptyBirth)

	err = usr.ChangeBirth(birth)

	assert.Exactly(t, birth, usr.Birth())
	assert.Equal(t, birthTm, usr.Birth().Time())
	assert.NoError(t, err)

	newCountry := shared.Country("X")
	err = usr.ChangeCountry(newCountry)

	assert.ErrorIs(t, err, shared.ErrNotACountry)

	err = usr.ChangeCountry(country)

	assert.Exactly(t, country, usr.Country())
	assert.Equal(t, countryStr, usr.Country().String())
	assert.NoError(t, err)

	newLanguage := shared.Language("X")
	err = usr.ChangeLanguage(newLanguage)

	assert.ErrorIs(t, err, shared.ErrNotALanguage)

	err = usr.ChangeLanguage(language)

	assert.Exactly(t, language, usr.Language())
	assert.Equal(t, languageStr, usr.Language().String())
	assert.NoError(t, err)

	usr.ChangePhone(nil)

	assert.Nil(t, usr.Phone())

	usr.ChangePhone(phone)

	assert.Exactly(t, phone, usr.Phone())
	assert.Equal(t, phoneStr, usr.Phone().String())

	information := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed dignissim, mi sed pulvinar feugiat," +
		"nisl erat commodo lorem, nec tincidunt magna orci ut odio. Praesent ullamcorper feugiat diam, sit amet tincidunt" +
		"nulla porta at. Aenean malesuada, lorem quis tristique feugiat, velit ligula convallis erat, sed luctus elit ex" +
		"non justo. Donec blandit velit ac elit gravida, sit amet maximus justo congue. Morbi ac semper justo, et luctus" +
		"erat. Suspendisse eget erat ac eros luctus pretium. Integer feugiat lacus et luctus dictum. Vestibulum ante ipsum " +
		"primis in faucibus orci luctus et ultrices posuere cubilia curae; Cras fermentum, erat eget ullamcorper efficitur, " +
		"justo risus pretium magna, ut dignissim arcu felis ac nulla."

	err = usr.ChangeInformation(&information)

	assert.ErrorIs(t, err, ErrLongInformationUser)

	information = "Just normal information"

	err = usr.ChangeInformation(&information)

	assert.Equal(t, information, *usr.Information())
	assert.NoError(t, err)

	err = usr.ChangeInformation(nil)

	assert.Nil(t, usr.Information())
	assert.NoError(t, err)

	profilePic := "./images/id/pic.jpg"
	usr.ChangeProfilePic(&profilePic)

	assert.Equal(t, profilePic, *usr.ProfilePic())

	usr.ChangeProfilePic(nil)

	assert.Nil(t, usr.ProfilePic())

	webSiteStr := "https://www.github.com/carlosclavijo/"
	newWebSite, err := shared.NewWebSite(&webSiteStr)

	assert.NotNil(t, newWebSite)
	assert.NoError(t, err)

	usr.ChangeWebSite(newWebSite)

	assert.Equal(t, newWebSite, usr.WebSite())

	usr.ChangeWebSite(nil)

	assert.Nil(t, usr.WebSite())

	usr.ChangeVisibility(true)

	assert.True(t, usr.Visibility())

	usr.ChangeVisibility(false)

	assert.False(t, usr.Visibility())

	oldTime := usr.LastLoginAt()
	time.Sleep(10 * time.Millisecond)

	usr.ChangeLastLoginAt()

	assert.True(t, usr.LastLoginAt().After(oldTime))

	oldTime = usr.UpdatedAt()
	time.Sleep(10 * time.Millisecond)

	usr.Update()

	assert.True(t, usr.UpdatedAt().After(oldTime))

	err = usr.Delete()

	assert.NotNil(t, usr.DeletedAt())
	assert.WithinDuration(t, time.Now(), *usr.DeletedAt(), time.Second)
	assert.NoError(t, err)

	err = usr.Delete()

	assert.ErrorIs(t, err, ErrAlreadyDeletedUser)

	err = usr.Restore()

	assert.Nil(t, usr.DeletedAt())
	assert.NoError(t, err)

	err = usr.Restore()

	assert.ErrorIs(t, err, ErrAlreadyRestoredUser)
}

func TestNewUserFromDB(t *testing.T) {
	id := uuid.New()
	firstName := "John"
	lastName := "Doe"
	userName := ""
	email := ""
	password := ""
	gender := "X"
	birth := time.Now()
	country := ""
	language := ""
	phone := "-"
	information := "A lot of information"
	profilePic := "./images/user/random-id.jpg"
	website := "-"
	lastLoginAt := time.Now()
	createdAt := time.Now()
	updatedAt := time.Now()
	deletedAt := time.Now().AddDate(0, 0, 1)

	usr, err := NewUserFromDB(id, firstName, lastName, userName, email, password, gender, birth, country, language, &phone, &information, &profilePic, &website, true, lastLoginAt, createdAt, updatedAt, &deletedAt)
	assert.Nil(t, usr)
	assert.ErrorIs(t, err, shared.ErrEmptyUsername)

	userName = "carlosclavijo"
	usr, err = NewUserFromDB(id, firstName, lastName, userName, email, password, gender, birth, country, language, &phone, &information, &profilePic, &website, true, lastLoginAt, createdAt, updatedAt, &deletedAt)
	assert.Nil(t, usr)
	assert.ErrorIs(t, err, shared.ErrEmptyEmail)

	email = "john@doe.com"
	usr, err = NewUserFromDB(id, firstName, lastName, userName, email, password, gender, birth, country, language, &phone, &information, &profilePic, &website, true, lastLoginAt, createdAt, updatedAt, &deletedAt)
	assert.Nil(t, usr)
	assert.Error(t, shared.ErrEmptyPassword)

	password = "5Trong!."
	usr, err = NewUserFromDB(id, firstName, lastName, userName, email, password, gender, birth, country, language, &phone, &information, &profilePic, &website, true, lastLoginAt, createdAt, updatedAt, &deletedAt)
	assert.Nil(t, usr)
	assert.ErrorIs(t, err, shared.ErrNotAGender)

	gender = "Male"
	usr, err = NewUserFromDB(id, firstName, lastName, userName, email, password, gender, birth, country, language, &phone, &information, &profilePic, &website, true, lastLoginAt, createdAt, updatedAt, &deletedAt)
	assert.Nil(t, usr)
	assert.ErrorIs(t, err, shared.ErrUnderTwelve)

	birth = time.Now().AddDate(-20, 0, 1)
	usr, err = NewUserFromDB(id, firstName, lastName, userName, email, password, gender, birth, country, language, &phone, &information, &profilePic, &website, true, lastLoginAt, createdAt, updatedAt, &deletedAt)
	assert.Nil(t, usr)
	assert.ErrorIs(t, err, shared.ErrNotACountry)

	country = "Bolivia"
	usr, err = NewUserFromDB(id, firstName, lastName, userName, email, password, gender, birth, country, language, &phone, &information, &profilePic, &website, true, lastLoginAt, createdAt, updatedAt, &deletedAt)
	assert.Nil(t, usr)
	assert.ErrorIs(t, err, shared.ErrNotALanguage)

	language = "Spanish"
	usr, err = NewUserFromDB(id, firstName, lastName, userName, email, password, gender, birth, country, language, &phone, &information, &profilePic, &website, true, lastLoginAt, createdAt, updatedAt, &deletedAt)
	assert.Nil(t, usr)
	assert.ErrorIs(t, err, shared.ErrNotNumericPhoneNumber)

	phone = "+591-70926048"
	usr, err = NewUserFromDB(id, firstName, lastName, userName, email, password, gender, birth, country, language, &phone, &information, &profilePic, &website, true, lastLoginAt, createdAt, updatedAt, &deletedAt)
	assert.Nil(t, usr)
	assert.ErrorIs(t, err, shared.ErrInvalidWebsite)

	website = "https://www.randomwebsite.com/profile?v=id"
	usr, err = NewUserFromDB(id, firstName, lastName, userName, email, password, gender, birth, country, language, &phone, &information, &profilePic, &website, true, lastLoginAt, createdAt, updatedAt, &deletedAt)

	require.NotNil(t, usr)
	require.NoError(t, err)
	assert.NotNil(t, usr.Id())
	assert.Equal(t, firstName, usr.FirstName())
	assert.Equal(t, lastName, usr.LastName())
	assert.Equal(t, userName, usr.Username().String())
	assert.Equal(t, email, usr.Email().String())
	assert.Equal(t, password, usr.Password().String())
	assert.Equal(t, gender, usr.Gender().String())
	assert.Equal(t, birth, usr.Birth().Time())
	assert.Equal(t, country, usr.Country().String())
	assert.Equal(t, language, usr.Language().String())
	assert.Equal(t, phone, usr.Phone().String())
	assert.Equal(t, information, *usr.Information())
	assert.Equal(t, profilePic, *usr.ProfilePic())
	assert.Equal(t, website, usr.WebSite().String())
	assert.True(t, usr.Visibility())
	assert.Equal(t, lastLoginAt, usr.LastLoginAt())
	assert.Equal(t, createdAt, usr.CreatedAt())
	assert.Equal(t, updatedAt, usr.UpdatedAt())
	assert.Equal(t, deletedAt, *usr.DeletedAt())
}
