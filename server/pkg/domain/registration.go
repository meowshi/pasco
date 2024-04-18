package domain

import "github.com/google/uuid"

// здесь мы добавим поле friends
type Registration struct {
	EventUuid      uuid.UUID `json:"event_uuid" db:"event_uuid"`
	YandexoidLogin string    `json:"yandexoid_login" db:"yandexoid_login"`
	Friends        int       `json:"friends" db:"friends"`
	Status         int       `json:"status" db:"status"`
	StatusCell     string    `json:"status_cell" db:"status_cell"`
}
