package usecase

import (
	"github.com/google/uuid"
	"github.com/meowshi/pasco-server/pkg/domain"
)

type PickUsecase interface {
	GetPickHistory() ([]*domain.Pick, error)
	CreatePick(req *domain.CreatePickReq) (int64, error)
	UpdateFromList(id int64, eventUuid uuid.UUID, peopleCount int, isListSuccess bool) error
	UpdateFromGift(id int64, isGiftSuccess bool) error
	UpdateFromBracelet(id int64, isBraceletSuccess bool) error
}
