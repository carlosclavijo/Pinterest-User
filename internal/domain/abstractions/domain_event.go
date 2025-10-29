package abstractions

import (
	"github.com/google/uuid"
	"time"
)

type DomainEvent struct {
	id         uuid.UUID
	occurredOn time.Time
}

func NewDomainEvent() *DomainEvent {
	return &DomainEvent{
		id:         uuid.New(),
		occurredOn: time.Now(),
	}
}

func (domain *DomainEvent) Id() uuid.UUID {
	return domain.id
}

func (domain *DomainEvent) OccurredOn() time.Time {
	return domain.occurredOn
}
