package usecase

import (
	"github.com/google/uuid"
	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/meowshi/pasco-server/pkg/storage"
)

type pickUsecase struct {
	storage storage.PickStorage
}

func NewPickUsecase(storage storage.PickStorage) *pickUsecase {
	return &pickUsecase{
		storage: storage,
	}
}

func (u *pickUsecase) GetPickHistory() ([]*domain.Pick, error) {
	picks, err := u.storage.GetPickHistory()
	if err != nil {
		return nil, err
	}

	return picks, nil
}

func (u *pickUsecase) CreatePick(req *domain.CreatePickReq) (int64, error) {
	id, err := u.storage.CreatePick(req)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *pickUsecase) UpdateFromList(id int64, eventUuid uuid.UUID, peopleCount int, isListSuccess bool) error {
	withFriends := !(peopleCount == 1)
	err := u.storage.UpdateFromList(id, eventUuid, withFriends, isListSuccess)
	if err != nil {
		return err
	}

	return nil
}

func (u *pickUsecase) UpdateFromGift(id int64, isGiftSuccess bool) error {
	err := u.storage.UpdateFromGift(id, isGiftSuccess)
	if err != nil {
		return err
	}

	return nil
}

func (u *pickUsecase) UpdateFromBracelet(id int64, isBraceletSuccess bool) error {
	err := u.storage.UpdateFromBracelet(id, isBraceletSuccess)
	if err != nil {
		return err
	}

	return nil
}
