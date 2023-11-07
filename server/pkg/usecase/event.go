package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/meowshi/pasco-server/pkg/domain"
)

type EventUsecase interface {
	Create(eventTitleGoogleSheetCell string) error
	Delete(eventUuid uuid.UUID) error
	Get(eventUuid uuid.UUID) (*domain.Event, error)
	GetByDate(date time.Time) ([]*domain.Event, error)
	GetEventLists(eventUuid uuid.UUID) ([]*domain.ListEntry, error)
	Update(event *domain.Event) error
}
