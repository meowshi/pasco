package usecase

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/meowshi/pasco-server/pkg/storage"
	"github.com/meowshi/pasco-server/pkg/types"
	"github.com/meowshi/pasco-server/pkg/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/sheets/v4"
)

type eventUsecase struct {
	service *sheets.Service

	eventStorage        storage.EventStorage
	registrationStorage storage.RegistrationStorage
	yandexoidStorage    storage.YandexoidStorage
	envStorage          storage.EnvStorage
}

func NewEventUsecase(service *sheets.Service, eventStorage storage.EventStorage, registrationStorage storage.RegistrationStorage, yandexoidStorage storage.YandexoidStorage, envStorage storage.EnvStorage) *eventUsecase {
	return &eventUsecase{
		service:             service,
		eventStorage:        eventStorage,
		registrationStorage: registrationStorage,
		yandexoidStorage:    yandexoidStorage,
		envStorage:          envStorage,
	}
}

// проверять обновлось ли название таблицы только здесь при создании нового мероприятия
func (u *eventUsecase) Create(eventTitleGoogleSheetCell string) error {
	spreadsheetId, err := u.envStorage.Get("GOOGLE_SPREADSHEET_ID")
	if err != nil {
		return err
	}

	sheetTitle, err := u.envStorage.Get("GOOGLE_SHEET_TITLE")
	if err != nil {
		return err
	}

	// Обновляем sheet title потому что его меняют каждый день
	if sheetTitle.EditedAt.Day() != time.Now().Day() {
		sheetId, err := u.envStorage.Get("GOOGLE_SHEET_ID")
		if err != nil {
			logrus.Errorf("Ошибка при получении GOOGLE_SHEET_ID из БД: %s.", err)
			return err
		}

		spreadsheet, err := u.service.Spreadsheets.Get(spreadsheetId.Value).Do()
		if err != nil {
			return err
		}

		sheetIdInt, _ := strconv.Atoi(sheetId.Value)
		for _, sheet := range spreadsheet.Sheets {
			if sheet.Properties.SheetId == int64(sheetIdInt) {
				sheetTitle.Value = sheet.Properties.Title
				sheetTitle.EditedAt = time.Now()
			}
		}

		err = u.envStorage.Update(sheetTitle)
		if err != nil {
			return err
		}
	}

	cell := types.NewCellFromString(eventTitleGoogleSheetCell)
	rng := fmt.Sprintf("%s!%s:%s", sheetTitle.Value, cell.ToString(), cell.AddColumn(3).Column)

	res, err := u.service.Spreadsheets.Values.Get(spreadsheetId.Value, rng).Do()
	if err != nil {
		logrus.Errorf("Ошибка при получении значений из google sheet. Range: %s. err: %s", rng, err)
		return err
	}

	values := res.Values
	if len(values[0][0].(string)) == 0 || len(values) < 2 {
		return errors.New("hазмер полученный данных из список меньше необходимого, видимо, ячейка указана неверно или списки дефектные")
	}

	eventName := values[0][0].(string)
	cuttedEventName := utils.CutEventName(eventName)

	startCell := types.NewCellFromString(eventTitleGoogleSheetCell)
	isInPlusOne, _ := u.eventStorage.CheckPlusOneEvent(cuttedEventName)

	event := &domain.Event{
		Uuid:            uuid.New(),
		Name:            eventName,
		GoogleSheetCell: eventTitleGoogleSheetCell,
		LockerEventId:   -1,
		CreatedAt:       time.Now(),
		AllowedFriends:  isInPlusOne,
	}

	yandexoids := make([]*domain.Yandexoid, 0)
	registrations := make([]*domain.Registration, 0)

	startCell.AddRow(1).AddColumn(3)
	people := values[1:]
	for _, val := range people {
		if len(val) < 3 || utils.HaveEmpty(val) {
			break
		}

		var status int
		statusCell := startCell.ToString()
		if len(val) == 4 {
			status, _ = strconv.Atoi(val[3].(string))
		} else {
			status = 0
		}

		yandexoid := &domain.Yandexoid{
			Login:   val[0].(string),
			Name:    val[1].(string),
			Surname: val[2].(string),
		}
		yandexoids = append(yandexoids, yandexoid)

		registrations = append(registrations, &domain.Registration{
			EventUuid:      event.Uuid,
			YandexoidLogin: yandexoid.Login,
			Status:         status,
			StatusCell:     statusCell,
		})

		startCell.AddRow(1)
	}

	err = u.eventStorage.Create(event)
	if err != nil {
		logrus.Errorf("Ошибка при создании события в БД: %s.", err)
		return err
	}

	err = u.yandexoidStorage.CreateMultiple(yandexoids...)
	if err != nil {
		logrus.Errorf("Ошибка при создании яндексоидов в БД: %s.", err)
		return err
	}

	err = u.registrationStorage.CreateMultiple(registrations...)
	if err != nil {
		logrus.Errorf("Ошибка при создании регистраций в БД: %s.", err)
		return err
	}

	return nil
}

func (u *eventUsecase) Delete(eventUuid uuid.UUID) error {
	err := u.eventStorage.Delete(eventUuid)

	return err
}

func (u *eventUsecase) Get(eventUuid uuid.UUID) (*domain.Event, error) {
	logrus.Fatalf("event get not implemented")

	return nil, errors.New("Not impl.")
}

func (u *eventUsecase) GetByDate(date time.Time) ([]*domain.Event, error) {
	events, err := u.eventStorage.GetByDate(date)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (u *eventUsecase) Update(event *domain.Event) error {
	err := u.eventStorage.Update(event)
	if err != nil {
		return err
	}

	cuttedEventName := utils.CutEventName(event.Name)
	if event.AllowedFriends {
		u.eventStorage.CreatePlusOneEvent(cuttedEventName)
	} else {
		u.eventStorage.DeletePlusOneEvent(cuttedEventName)
	}

	return nil
}

func (u *eventUsecase) GetEventLists(eventUuid uuid.UUID) ([]*domain.ListEntry, error) {
	list, err := u.eventStorage.GetEventLists(eventUuid)
	if err != nil {
		return nil, err
	}

	return list, nil
}
