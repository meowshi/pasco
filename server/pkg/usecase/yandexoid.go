package usecase

import (
	"time"

	"github.com/meowshi/pasco-server/pkg/domain"
)

type YandexoidUsecase interface {
	Get(yandexoidLogin string) (*domain.Yandexoid, error)
	GetEventsByDate(yandexoidLogin string, date time.Time) ([]*domain.EventWithYandexoidStatusCell, error)
	GetEvents(yandexoidLogin string) ([]*domain.EventWithYandexoidStatusCell, error)
	Create(yandexoid *domain.Yandexoid) error
}
