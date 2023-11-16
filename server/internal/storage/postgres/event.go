package postgres

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/sirupsen/logrus"
)

type postgresEventStorage struct {
	db *sqlx.DB
}

func NewEventPostgresStorage(db *sqlx.DB) *postgresEventStorage {
	return &postgresEventStorage{
		db: db,
	}
}

func (s *postgresEventStorage) Create(event *domain.Event) error {
	result, err := s.db.Exec(
		`INSERT INTO event
		SELECT $1, $2, $3, $4, $5, $6
		WHERE NOT EXISTS(
			SELECT 1 from event WHERE name=$2 AND google_sheet_cell=$3 AND date_trunc('day', created_at::timestamptz)=date_trunc('day', $5::timestamptz)
		)`,
		event.Uuid,
		event.Name,
		event.GoogleSheetCell,
		event.LockerEventId,
		event.CreatedAt,
		event.AllowedFriends,
	)
	if err != nil {
		logrus.Errorf("Ошибка при создании события: %s.", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		err = errors.New("DB: event already exists.")
		logrus.Errorf("Ошибка при создании события: %s.", err)
		return err
	}

	return nil
}

func (s *postgresEventStorage) Delete(eventUuid uuid.UUID) error {
	_, err := s.db.Exec("DELETE FROM event WHERE uuid=$1", eventUuid)
	if err != nil {
		logrus.Errorf("Ошибка при удалении события: %s.", err)
		return err
	}

	return nil
}

func (s *postgresEventStorage) Update(event *domain.Event) error {
	_, err := s.db.Exec("UPDATE event SET name=$1, google_sheet_cell=$2, locker_event_id=$3, created_at=$4, allowed_friends=$5 WHERE uuid=$6",
		event.Name,
		event.GoogleSheetCell,
		event.LockerEventId,
		event.CreatedAt,
		event.AllowedFriends,
		event.Uuid,
	)
	if err != nil {
		logrus.Errorf("Ошибка при обновлении события: %s.", err)
		return err
	}

	return nil
}

func (s *postgresEventStorage) Get(eventUuid uuid.UUID) (*domain.Event, error) {
	event := &domain.Event{}

	err := s.db.Get(event, "SELECT * FROM event WHERE uuid=$1", eventUuid)
	if err != nil {
		logrus.Errorf("Ошибка при получении события по uuid: %s.", err)
		return nil, err
	}

	return event, nil
}

func (s *postgresEventStorage) GetByDate(date time.Time) ([]*domain.Event, error) {
	var events []*domain.Event

	err := s.db.Select(&events,
		`SELECT * FROM event
		WHERE date_trunc('day', created_at::timestamptz)=date_trunc('day', $1::timestamptz)
		ORDER BY name
		`,
		date,
	)
	if err != nil {
		logrus.Errorf("Ошибка при получении событий по дате: %s.", err)
		return nil, err
	}

	return events, nil
}

func (s *postgresEventStorage) GetEventLists(eventUuid uuid.UUID) ([]*domain.ListEntry, error) {
	var list []*domain.ListEntry

	err := s.db.Select(&list,
		`SELECT yandexoid.*, registration.status
		FROM registration
		JOIN yandexoid ON registration.yandexoid_login=yandexoid.login
		WHERE registration.event_uuid=$1
		ORDER BY yandexoid.login`,
		eventUuid,
	)
	if err != nil {
		logrus.Errorf("Ошибка при получении списка яндексоидов на мепроприятие: %s.", err)
		return nil, err
	}

	return list, nil
}

func (s *postgresEventStorage) CreatePlusOneEvent(name string) error {
	_, err := s.db.Exec(`
		INSERT INTO plus_one_event (name)
		VALUES ($1)
		ON CONFLICT (name) DO NOTHING`,
		name,
	)
	if err != nil {
		logrus.Errorf("Ошибка при создании события+1: %s.", err)
		return err
	}

	return nil
}

func (s *postgresEventStorage) DeletePlusOneEvent(name string) error {
	_, err := s.db.Exec(`
		DELETE FROM plus_one_event
		WHERE name=$1`,
		name,
	)
	if err != nil {
		logrus.Errorf("Ошибка при удалении события+1: %s.", err)
		return err
	}

	return nil
}

func (s *postgresEventStorage) CheckPlusOneEvent(name string) (bool, error) {
	row, err := s.db.Query(`
		SELECT EXISTS(SELECT 1 FROM plus_one_event WHERE name=$1)`,
		name,
	)
	if err != nil {
		logrus.Errorf("Ошибка при проверке существования события+1: %s.", err)
		return false, err
	}
	exists := false

	row.Next()
	row.Scan(&exists)

	err = row.Close()
	if err != nil {
		return false, err
	}

	return exists, nil
}
