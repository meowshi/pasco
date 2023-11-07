package usecase

import (
	"time"

	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/meowshi/pasco-server/pkg/storage"
)

type yandexoidUsecase struct {
	yandexoidStorage storage.YandexoidStorage
}

func NewYandexoidUsecase(ys storage.YandexoidStorage) *yandexoidUsecase {
	return &yandexoidUsecase{
		yandexoidStorage: ys,
	}
}

func (u *yandexoidUsecase) Get(yandexoidLogin string) (*domain.Yandexoid, error) {
	y, err := u.yandexoidStorage.Get(yandexoidLogin)
	if err != nil {
		return nil, err
	}

	return y, nil
}

func (u *yandexoidUsecase) GetEventsByDate(yandexoidLogin string, date time.Time) ([]*domain.EventWithYandexoidStatusCell, error) {
	events, err := u.yandexoidStorage.GetEventsByDate(yandexoidLogin, date)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (u *yandexoidUsecase) GetEvents(yandexoidLogin string) ([]*domain.EventWithYandexoidStatusCell, error) {
	events, err := u.yandexoidStorage.GetEvents(yandexoidLogin)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (u *yandexoidUsecase) Create(yandexoid *domain.Yandexoid) error {
	err := u.yandexoidStorage.Create(yandexoid)
	if err != nil {
		return err
	}

	return nil
}
