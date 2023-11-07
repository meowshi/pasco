package storage

import (
	"time"

	"github.com/google/uuid"
	"github.com/meowshi/pasco-server/pkg/domain"
)

type EventStorage interface {
	Create(event *domain.Event) error
	Delete(eventUuid uuid.UUID) error
	Update(event *domain.Event) error
	Get(eventUuid uuid.UUID) (*domain.Event, error)
	GetByDate(date time.Time) ([]*domain.Event, error)
	GetEventLists(eventUuid uuid.UUID) ([]*domain.ListEntry, error)
	CreatePlusOneEvent(name string) error
	DeletePlusOneEvent(name string) error
	CheckPlusOneEvent(name string) (bool, error)
}
