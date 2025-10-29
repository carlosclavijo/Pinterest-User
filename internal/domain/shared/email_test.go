package shared

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewEmail(t *testing.T) {
	name := "valid@email.com"

	email, err := NewEmail(name)

	require.NotEmpty(t, email)
	require.NoError(t, err)
	assert.NotEmpty(t, email.Local())
	assert.NotEmpty(t, email.Domain())
	assert.Equal(t, name, email.String())
	assert.Equal(t, "valid", email.Local())
	assert.Equal(t, "email.com", email.Domain())
}

func TestEmail_Empty(t *testing.T) {
	empty := ""

	email, err := NewEmail(empty)

	require.NotNil(t, email)
	require.Empty(t, email)
	assert.ErrorIs(t, err, ErrEmptyEmail)
}

func TestEmail_Invalid(t *testing.T) {
	invalid := "invalidemail.com"

	email, err := NewEmail(invalid)

	require.NotNil(t, email)
	require.Empty(t, email)
	assert.ErrorIs(t, err, ErrInvalidEmail)
}

func TestEmail_LongLocal(t *testing.T) {
	name := "longlocalnametoolongthatcannotbealocalnamesowouldfailedeveryunittest@test.com"

	email, err := NewEmail(name)

	require.NotNil(t, email)
	require.Empty(t, email)
	assert.ErrorIs(t, err, ErrLongLocalEmail)
}

func TestEmail_LongDomain(t *testing.T) {
	name := "test@toomuchlongdomainnamethatwouldfailedeveryoneofthetesteventhehardestoneandidontknowwhatelsetoputheretoreachthe255oflongandevenmoretoreachthefailmaybeishouldtrytousechatgptforthisonebutiamtiredofthattoolnotbecauseisnnthelpfulbutbecauseaiistakingmyjobandinneedtofindajobtolivethatswhyiamdoingtheselivestreamspleaseifyouarereadingthisgivemeajobineedthemoneyipromiseiwillpushmylimitstowritegoodandqualityjobs.com"

	email, err := NewEmail(name)

	require.NotNil(t, email)
	require.Empty(t, email)
	assert.ErrorIs(t, err, ErrLongDomainEmail)
}
