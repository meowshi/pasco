package domain

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Uuid            uuid.UUID `db:"uuid" json:"uuid"`
	Name            string    `db:"name" json:"name"`
	GoogleSheetCell string    `db:"google_sheet_cell" json:"google_sheet_cell"`
	LockerEventId   int       `db:"locker_event_id" json:"locker_event_id"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	AllowedFriends  bool      `db:"allowed_friends" json:"allowed_friends"`
}
