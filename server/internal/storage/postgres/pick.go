package postgres

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/sirupsen/logrus"
)

type pickPostgresStorage struct {
	db *sqlx.DB
}

func NewPickPostgresStorage(db *sqlx.DB) *pickPostgresStorage {
	return &pickPostgresStorage{
		db: db,
	}
}

func (s *pickPostgresStorage) GetPickHistory() ([]*domain.Pick, error) {
	var picks []*domain.Pick
	err := s.db.Select(&picks,
		`SELECT yandexoid.*, event.name AS event_name, pick.with_friends, pick.is_list_success, pick.is_gift_success, pick.is_bracelet_success, pick.picked_at FROM pick
		JOIN yandexoid ON yandexoid.login=pick.yandexoid_login
		JOIN event ON event.uuid=pick.event_uuid
		ORDER BY id DESC
		LIMIT 10 
		`,
	)
	if err != nil {
		logrus.Errorf("Ошибка при получении истории пиков: %s.", err)
		return nil, err
	}

	return picks, nil
}

func (s *pickPostgresStorage) CreatePick(req *domain.CreatePickReq) (int64, error) {
	var id int64 = 0
	rows, err := s.db.NamedQuery(`
		INSERT INTO pick
		VALUES (:yandexoid_login, :event_uuid, :with_friends, :is_list_success, :is_gift_success, :is_bracelet_success, :picked_at, DEFAULT)
		RETURNING id`,
		req,
	)
	if err != nil {
		logrus.Errorf("Ошибка при создании пика: %s.", err)
		return 0, err
	}

	rows.Next()
	err = rows.Scan(&id)
	if err != nil {
		logrus.Errorf("Ошибка при сканировании возвращенного pick id: %s.", err)
		return 0, err
	}

	err = rows.Close()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *pickPostgresStorage) UpdateFromList(id int64, eventUuid uuid.UUID, withFriends, isListSuccess bool) error {
	_, err := s.db.Exec(`
		UPDATE pick
		SET event_uuid=$1, with_friends=$2, is_list_success=$3
		WHERE id=$4`,
		eventUuid,
		withFriends,
		isListSuccess,
		id,
	)
	if err != nil {
		logrus.Errorf("Ошибка при обновлении пика из списков: %s.", err)
		return err
	}

	return nil
}

func (s *pickPostgresStorage) UpdateFromGift(id int64, isGiftSuccess bool) error {
	_, err := s.db.Exec(`
		UPDATE pick
		SET is_gift_success=$1
		WHERE id=$2`,
		isGiftSuccess,
		id,
	)
	if err != nil {
		logrus.Errorf("Ошибка при обновлении пика из подарков: %s.", err)
		return err
	}

	return nil
}

func (s *pickPostgresStorage) UpdateFromBracelet(id int64, isBraceletSuccess bool) error {
	_, err := s.db.Exec(`
		UPDATE pick
		SET is_bracelet_success=$1
		WHERE id=$2`,
		isBraceletSuccess,
		id,
	)
	if err != nil {
		logrus.Errorf("Ошибка при обновлении пика из браслетов: %s.", err)
		return err
	}

	return nil
}
