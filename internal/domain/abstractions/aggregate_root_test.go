package abstractions

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewAggregateRoot(t *testing.T) {
	id := uuid.New()

	aggregateRoot := NewAggregateRoot(id)

	require.NotEmpty(t, aggregateRoot)
	assert.NotNil(t, aggregateRoot.Id)
}
