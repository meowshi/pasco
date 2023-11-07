package storage

import (
	"github.com/google/uuid"
	"github.com/meowshi/pasco-server/pkg/domain"
)

type PickStorage interface {
	GetPickHistory() ([]*domain.Pick, error)
	CreatePick(req *domain.CreatePickReq) (int64, error)
	UpdateFromList(id int64, eventUuid uuid.UUID, withFriends, isListSuccess bool) error
	UpdateFromGift(id int64, isGiftSuccess bool) error
	UpdateFromBracelet(id int64, isBraceletSuccess bool) error
}
