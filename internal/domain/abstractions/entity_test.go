package abstractions

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewEntity(t *testing.T) {
	id := uuid.New()

	entity := NewEntity(id)

	require.NotEmpty(t, entity)
	assert.NotNil(t, entity.Id)
	assert.Empty(t, entity.DomainEvents())

	domainEvent1 := NewDomainEvent()
	domainEvent2 := NewDomainEvent()
	domainEvent3 := NewDomainEvent()

	entity.AddDomainEvent(*domainEvent1)
	entity.AddDomainEvent(*domainEvent2)
	entity.AddDomainEvent(*domainEvent3)

	assert.NotEmpty(t, entity.DomainEvents())
	assert.Len(t, entity.DomainEvents(), 3)

	entity.ClearDomainEvents()

	assert.Empty(t, entity.DomainEvents())
	assert.Len(t, entity.DomainEvents(), 0)
}
