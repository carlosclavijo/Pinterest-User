package abstractions

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDomainEvent(t *testing.T) {
	domainEvent := NewDomainEvent()

	require.NotEmpty(t, domainEvent)
	assert.NotNil(t, domainEvent.Id())
	assert.NotNil(t, domainEvent.OccurredOn())
}
