package storage

import (
	"time"

	"github.com/meowshi/pasco-server/pkg/domain"
)

type YandexoidStorage interface {
	Create(yandexoid *domain.Yandexoid) error
	CreateMultiple(yandexoids ...*domain.Yandexoid) error
	Delete(yandexoidLogin string) error
	Get(yandexoidLogin string) (*domain.Yandexoid, error)
	GetEventsByDate(yandexoidLogin string, date time.Time) ([]*domain.EventWithYandexoidStatusCell, error)
	GetEvents(yandexoidLogin string) ([]*domain.EventWithYandexoidStatusCell, error)
}
