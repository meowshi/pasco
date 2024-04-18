package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/sirupsen/logrus"
)

type lockerUsecase struct {
	client *http.Client
}

func NewLockerUsecase() *lockerUsecase {
	return &lockerUsecase{
		client: &http.Client{},
	}
}

func (u *lockerUsecase) GetLockerEvents() ([]*domain.LockerEvent, error) {
	date := time.Now()
	req, err := http.NewRequest("GET", os.Getenv("LOCKER_API")+"/events/events/", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("start_time__gte", fmt.Sprintf("%d-%.2d-%dT00:00:00Z", date.Year(), date.Month(), date.Day()))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", os.Getenv("LOCKER_TOKEN"))

	res, err := u.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var lockerEvents []*domain.LockerEvent
	err = json.Unmarshal(body, &lockerEvents)
	if err != nil {
		return nil, err
	}

	return lockerEvents, err
}

func (u *lockerUsecase) PrintBracelet(printReq *domain.PrintBraceletReq) error {
	body, err := json.Marshal(printReq)
	if err != nil {
		logrus.Errorf("Ошибка при unmarshall PrintBraceletReq: %s.", err)
		return err
	}

	req, err := http.NewRequest("POST", os.Getenv("LOCKER_API")+"/bracelets/create/", bytes.NewReader(body))
	if err != nil {
		logrus.Errorf("Ошибка при формировании запроса: %s.", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Authorization", os.Getenv("LOCKER_TOKEN"))

	res, err := u.client.Do(req)
	if err != nil {
		logrus.Errorf("Ошибка при выполнении запроса на печать браслетов: %s.", err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = fmt.Errorf("status code is not OK. %s", res.Status)
		logrus.Errorf("Ошибка при выполнении запроса на печать браслетов: %s.", err)
		return err
	}

	return nil
}

func (u *lockerUsecase) GetPrinters() ([]*domain.Printer, error) {
	req, err := http.NewRequest("GET", os.Getenv("LOCKER_API")+"/workstation/bracelet-printers/", nil)
	if err != nil {
		logrus.Errorf("Ошибка при создании запроса на получение принтеров: %s.", err)
		return nil, err
	}

	req.Header.Set("Authorization", os.Getenv("LOCKER_TOKEN"))

	res, err := u.client.Do(req)
	if err != nil {
		logrus.Errorf("Ошибка при выполнении запроса на получение принтеров: %s", err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != fiber.StatusOK {
		err = fmt.Errorf("status code is not OK. %s", res.Status)
		logrus.Errorf("Ошибка при выполнении запроса на получение принтеров: %s.", err)
		return nil, err
	}

	printers := make([]*domain.Printer, 0)
	err = json.NewDecoder(res.Body).Decode(&printers)
	if err != nil {
		logrus.Errorf("Ошибка при парсинге тела GetPrinters: %s.", err)
		return nil, err
	}

	return printers, nil
}
