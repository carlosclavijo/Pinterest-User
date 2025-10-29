package abstractions

import "github.com/google/uuid"

type Entity struct {
	Id           uuid.UUID
	domainEvents []DomainEvent
}

func NewEntity(id uuid.UUID) *Entity {
	return &Entity{
		Id:           id,
		domainEvents: []DomainEvent{},
	}
}

func (entity *Entity) DomainEvents() []DomainEvent {
	return entity.domainEvents
}

func (entity *Entity) AddDomainEvent(event DomainEvent) {
	entity.domainEvents = append(entity.domainEvents, event)
}

func (entity *Entity) ClearDomainEvents() {
	entity.domainEvents = nil
}
