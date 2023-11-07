package domain

import (
	"time"

	"github.com/google/uuid"
)

type Pick struct {
	*Yandexoid
	EventName         string    `json:"event_name" db:"event_name"`
	WithFriends       bool      `json:"with_friends" db:"with_friends"`
	IsListSuccess     bool      `json:"is_list_success" db:"is_list_success"`
	IsGiftSuccess     bool      `json:"is_gift_success" db:"is_gift_success"`
	IsBraceletSuccess bool      `json:"is_bracelet_success" db:"is_bracelet_success"`
	PickedAt          time.Time `json:"picked_at" db:"picked_at"`
}

type CreatePickReq struct {
	YandexoidLogin    string    `json:"yandexoid_login" db:"yandexoid_login"`
	EventUuid         uuid.UUID `json:"event_uuid" db:"event_uuid"`
	WithFriends       bool      `json:"with_friends" db:"with_friends"`
	IsListSuccess     bool      `json:"is_list_success" db:"is_list_success"`
	IsGiftSuccess     bool      `json:"is_gift_success" db:"is_gift_success"`
	IsBraceletSuccess bool      `json:"is_bracelet_success" db:"is_bracelet_success"`
	PickedAt          time.Time `json:"picked_at" db:"picked_at"`
}
