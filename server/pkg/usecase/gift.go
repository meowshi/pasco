package usecase

import "github.com/meowshi/pasco-server/pkg/domain"

type GiftUsecase interface {
	Get(key string) (*domain.GetRes, error)
	Give(key, login, count string) error
}
