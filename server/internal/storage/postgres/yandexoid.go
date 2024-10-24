package postgres

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/sirupsen/logrus"
)

type yandexoidPostgresStorage struct {
	db *sqlx.DB
}

func NewYandexoidPostgresStorage(db *sqlx.DB) *yandexoidPostgresStorage {
	return &yandexoidPostgresStorage{
		db: db,
	}
}

func (s *yandexoidPostgresStorage) Create(yandexoid *domain.Yandexoid) error {
	_, err := s.db.NamedExec(`
		INSERT INTO yandexoid
		VALUES (:login, :name, :surname)
		ON CONFLICT (login)
		DO UPDATE
		SET name=excluded.name, surname=excluded.surname`,
		yandexoid)

	return err
}

func (s *yandexoidPostgresStorage) CreateMultiple(yandexoids ...*domain.Yandexoid) error {
	_, err := s.db.NamedExec(`
		INSERT INTO yandexoid (login, name, surname)
		VALUES (:login, :name, :surname)
		ON CONFLICT (login)
		DO UPDATE
		SET name=excluded.name, surname=excluded.surname`,
		yandexoids,
	)

	return err
}
func (s *yandexoidPostgresStorage) Delete(yandexoidLogin string) error {
	_, err := s.db.Exec("DELETE FROM yandexoid WHERE login=$1", yandexoidLogin)

	return err
}

func (s *yandexoidPostgresStorage) Get(yandexoidLogin string) (*domain.Yandexoid, error) {
	yandexoid := &domain.Yandexoid{}

	err := s.db.Get(yandexoid, "SELECT * FROM yandexoid WHERE login=$1", yandexoidLogin)
	if err != nil {
		return nil, err
	}

	return yandexoid, nil
}

func (s *yandexoidPostgresStorage) GetEventsByDate(yandexoidLogin string, date time.Time) ([]*domain.EventWithYandexoidStatusCell, error) {
	var events []*domain.EventWithYandexoidStatusCell

	err := s.db.Select(&events,
		`SELECT event.*, registration.friends, registration.status, registration.status_cell FROM event
		JOIN registration ON event.uuid=registration.event_uuid
		WHERE registration.yandexoid_login=$1 AND date_trunc('day', event.created_at::timestamptz)=date_trunc('day', $2::timestamptz)
		ORDER BY event.name`,
		yandexoidLogin,
		date,
	)
	if err != nil {
		logrus.Errorf("Ошибка при получении событий, на которые записан яндексоид в конктретный день: %s.", err)
		return nil, err
	}

	return events, nil
}

func (s *yandexoidPostgresStorage) GetEvents(yandexoidLogin string) ([]*domain.EventWithYandexoidStatusCell, error) {
	var events []*domain.EventWithYandexoidStatusCell

	err := s.db.Select(&events,
		`SELECT event.*, registration.friends, registration.status_cell FROM event
		JOIN registration ON event.uuid=registration.event_uuid
		WHERE registration.yandexoid_login=$1
		ORDER BY event.name
		LIMIT 50`,
		yandexoidLogin,
	)
	if err != nil {
		logrus.Errorf("Ошибка при получении событий, на которые записан яндексоид: %s.", err)
		return nil, err
	}

	return events, nil
}
