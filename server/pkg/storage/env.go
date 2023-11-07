package storage

import "github.com/meowshi/pasco-server/pkg/domain"

type EnvStorage interface {
	Get(envName string) (*domain.Env, error)
	Update(env *domain.Env) error
}
