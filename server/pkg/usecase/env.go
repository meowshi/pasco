package usecase

import "github.com/meowshi/pasco-server/pkg/domain"

type EnvUsecase interface {
	Get(key string) (*domain.Env, error)
	Update(env *domain.Env) error
	UpdateSheetTitle() error
}
