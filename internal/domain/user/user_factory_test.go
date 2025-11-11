package users

import (
	"github.com/carlosclavijo/Pinterest-Services/internal/domain/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewUserFactory(t *testing.T) {
	factory := NewUserFactory()

	require.NotNil(t, factory)
	require.Empty(t, factory)
}

func TestUserFactory_Create(t *testing.T) {
	factory := NewUserFactory()

	require.NotNil(t, factory)
	require.Empty(t, factory)

	firstName := ""
	lastName := ""
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

	usr, err := factory.Create(firstName, lastName, userName, email, password, gender, birth, country, language, phone)
	assert.Nil(t, usr)
	assert.ErrorIs(t, err, ErrEmptyFirstNameUser)

	firstName = "a"
	usr, err = factory.Create(firstName, lastName, userName, email, password, gender, birth, country, language, phone)
	assert.Nil(t, usr)
	assert.ErrorIs(t, err, ErrEmptyLastNameUser)

	firstName = "longnamethatdontsupposetobethislongasfirstnamesoineedtofillwithusefulinfotmationlikethisandsometimeswhenidothatireachwayhighercharactersnumbersthatiactuallyneeded"
	lastName = "longnamethatdontsupposetobethislongasfirstnamesoineedtofillwithusefulinfotmationlikethisandsometimeswhenidothatireachwayhighercharactersnumbersthatiactuallyneeded"
	usr, err = factory.Create(firstName, lastName, userName, email, password, gender, birth, country, language, phone)
	assert.Nil(t, usr)
	assert.ErrorIs(t, err, ErrLongFirstNameUser)

	firstName = "John5"
	usr, err = factory.Create(firstName, lastName, userName, email, password, gender, birth, country, language, phone)
	assert.Nil(t, usr)
	assert.ErrorIs(t, err, ErrLongLastNameUser)

	lastName = "Doe5"
	usr, err = factory.Create(firstName, lastName, userName, email, password, gender, birth, country, language, phone)
	assert.Nil(t, usr)
	assert.ErrorIs(t, err, ErrNonAlphaFirstNameUser)

	firstName = "John"
	usr, err = factory.Create(firstName, lastName, userName, email, password, gender, birth, country, language, phone)
	assert.Nil(t, usr)
	assert.ErrorIs(t, err, ErrNonAlphaLastNameUser)

	lastName = "Doe"

	usr, err = factory.Create(firstName, lastName, userName, email, password, gender, birth, country, language, phone)

	require.NotNil(t, usr)
	require.NoError(t, err)
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
}

func TestIsAlphaName(t *testing.T) {
	name := ""
	isAlpha := isAlphaName(name)

	name = "J0hn"
	isAlpha = isAlphaName(name)

	assert.False(t, isAlpha)

	name = "John"
	isAlpha = isAlphaName(name)

	assert.True(t, isAlpha)

	name = "John Doe"
	isAlpha = isAlphaName(name)

	assert.True(t, isAlpha)

	name = "John  Doe"
	isAlpha = isAlphaName(name)

	assert.False(t, isAlpha)
}
