package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/sirupsen/logrus"
)

type giftUsecase struct {
	client *http.Client
}

func NewGiftUsecase() *giftUsecase {
	return &giftUsecase{
		client: &http.Client{},
	}
}

func (u *giftUsecase) Get(key string) (*domain.GetRes, error) {
	req, err := http.NewRequest("GET", os.Getenv("GIFT_AUTH"), nil)
	if err != nil {
		logrus.Errorf("Ошибка при создании запроса gift auth: %s.", err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("key", key)
	req.URL.RawQuery = q.Encode()

	res, err := u.client.Do(req)
	if err != nil {
		logrus.Errorf("Ошибка при выполнении запроса gift auth: %s.", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("HTTP response status - %s", res.Status))
		logrus.Errorf("Gift auth response status code != 200: %s", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Errorf("Ошибка при чтении тела ответа gift auth: %s.", err)
		return nil, err
	}

	getRes := &domain.GetRes{}
	err = json.Unmarshal(body, &getRes)
	if err != nil {
		logrus.Errorf("Ошибка при unmarshall тела ответа gift auth: %s", err)
		return nil, err
	}

	return getRes, nil
}

func (u *giftUsecase) Give(key, login, count string) error {
	giftUrl := os.Getenv("GIFT_COLLECT") + login + "/"

	form := url.Values{}
	form.Add("key", key)
	form.Add("count", count)

	req, err := http.NewRequest(http.MethodPost, giftUrl, strings.NewReader(form.Encode()))
	if err != nil {
		logrus.Errorf("Ошибка при создании запроса gift collect: %s.", err)
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := u.client.Do(req)
	if err != nil {
		logrus.Errorf("Ошибка при выполнении запроса gift collect: %s.", err)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("Http response status - %s", res.Status))
		logrus.Errorf("Gift collect response status code != 200: %s.", err)
		return err
	}

	return nil
}
