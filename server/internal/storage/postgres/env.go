package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/meowshi/pasco-server/pkg/domain"
)

type envPostgresStorage struct {
	db *sqlx.DB
}

func NewEnvPostgresStorage(db *sqlx.DB) *envPostgresStorage {
	return &envPostgresStorage{
		db: db,
	}
}

func (s *envPostgresStorage) Get(envName string) (*domain.Env, error) {
	env := &domain.Env{}

	err := s.db.Get(env,
		`
			SELECT *
			FROM env
			WHERE key=$1
		`,
		envName,
	)

	if err != nil {
		return nil, err
	}

	return env, nil
}

func (s *envPostgresStorage) Update(env *domain.Env) error {
	_, err := s.db.NamedExec(`
		UPDATE env
		SET value=:value, edited_at=:edited_at
		WHERE key=:key`,
		env,
	)

	return err
}
