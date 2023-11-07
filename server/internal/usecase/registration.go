package usecase

import (
	"fmt"

	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/meowshi/pasco-server/pkg/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/sheets/v4"
)

type registrationUsecase struct {
	registrationStorage storage.RegistrationStorage
	envStorage          storage.EnvStorage

	service *sheets.Service
}

func NewRegistrationUsecase(rs storage.RegistrationStorage, es storage.EnvStorage, service *sheets.Service) *registrationUsecase {
	return &registrationUsecase{
		registrationStorage: rs,
		envStorage:          es,
		service:             service,
	}
}

func (u *registrationUsecase) Update(registration *domain.Registration) error {
	err := u.registrationStorage.Update(registration)
	if err != nil {
		return err
	}

	sheetTitle, err := u.envStorage.Get("GOOGLE_SHEET_TITLE")
	if err != nil {
		return err
	}

	spreadsheetId, err := u.envStorage.Get("GOOGLE_SPREADSHEET_ID")
	if err != nil {
		return err
	}

	statusRange := fmt.Sprintf("%s!%s", sheetTitle.Value, registration.StatusCell)
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, []interface{}{registration.Status})

	_, err = u.service.Spreadsheets.Values.Update(spreadsheetId.Value, statusRange, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		logrus.Errorf("Ошибка при изменении статуса в google sheet: %s.", err)
		return err
	}

	return nil
}
