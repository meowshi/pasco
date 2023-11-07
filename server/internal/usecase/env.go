package usecase

import (
	"strconv"
	"time"

	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/meowshi/pasco-server/pkg/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/sheets/v4"
)

type envUsecase struct {
	storage storage.EnvStorage
	service *sheets.Service
}

func NewEnvUsecase(storage storage.EnvStorage, service *sheets.Service) *envUsecase {
	return &envUsecase{
		storage: storage,
		service: service,
	}
}

func (u *envUsecase) Get(key string) (*domain.Env, error) {
	env, err := u.storage.Get(key)
	if err != nil {
		logrus.Errorf("Ошибка при получении переменной %s: %s.", key, err)
		return nil, err
	}

	return env, nil
}

func (u *envUsecase) Update(env *domain.Env) error {
	err := u.storage.Update(env)
	if err != nil {
		return err
	}

	return nil
}

func (u *envUsecase) UpdateSheetTitle() error {
	spreadsheetId, err := u.storage.Get("GOOGLE_SPREADSHEET_ID")
	if err != nil {
		return err
	}

	sheetId, err := u.storage.Get("GOOGLE_SHEET_ID")
	if err != nil {
		return err
	}

	spreadsheet, err := u.service.Spreadsheets.Get(spreadsheetId.Value).Do()
	if err != nil {
		return err
	}

	sheetTitle := &domain.Env{}
	sheetIdInt, _ := strconv.Atoi(sheetId.Value)
	for _, sheet := range spreadsheet.Sheets {
		if sheet.Properties.SheetId == int64(sheetIdInt) {
			sheetTitle.Key = "GOOGLE_SHEET_TITLE"
			sheetTitle.Value = sheet.Properties.Title
			sheetTitle.EditedAt = time.Now()
		}
	}

	err = u.storage.Update(sheetTitle)
	if err != nil {
		return err
	}

	return nil
}
