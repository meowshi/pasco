package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/sirupsen/logrus"
)

type registrationPostgresStorage struct {
	db *sqlx.DB
}

func NewRegistrationPostgresStorage(db *sqlx.DB) *registrationPostgresStorage {
	return &registrationPostgresStorage{
		db: db,
	}
}

func (s *registrationPostgresStorage) Create(registration *domain.Registration) error {
	_, err := s.db.NamedExec(`
		INSERT INTO registration
		VALUES (:event_uuid, :yandexoid_login, :friends, :status, :status_cell)
		ON CONFLICT(event_uuid, yandexoid_login)
		DO NOTHING`,
		registration,
	)

	return err
}

func (s *registrationPostgresStorage) CreateMultiple(registrations ...*domain.Registration) error {
	_, err := s.db.NamedExec(`
		INSERT INTO registration (event_uuid, yandexoid_login, friends, status, status_cell)
		VALUES (:event_uuid, :yandexoid_login, :friends, :status, :status_cell)`,
		registrations,
	)

	return err
}

func (s *registrationPostgresStorage) Update(registration *domain.Registration) error {
	_, err := s.db.NamedExec(
		`UPDATE registration
		SET status=:status
		WHERE event_uuid=:event_uuid AND yandexoid_login=:yandexoid_login
		`,
		registration,
	)
	if err != nil {
		logrus.Errorf("Ошибка при обновлении регистрации: %s.", err)
		return err
	}

	return nil
}
