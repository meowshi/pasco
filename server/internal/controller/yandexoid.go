package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/meowshi/pasco-server/pkg/usecase"
)

type YandexoidController struct {
	yandexoidUsecase usecase.YandexoidUsecase
	giftUsecase      usecase.GiftUsecase
	pickUsecase      usecase.PickUsecase
}

func NewYandexoidController(yu usecase.YandexoidUsecase, gu usecase.GiftUsecase, pu usecase.PickUsecase) *YandexoidController {
	return &YandexoidController{
		yandexoidUsecase: yu,
		giftUsecase:      gu,
		pickUsecase:      pu,
	}
}

func (c *YandexoidController) GetEvents(ctx *fiber.Ctx) error {
	login := ctx.Params("login")
	dateString := ctx.Query("date")

	var events []*domain.EventWithYandexoidStatusCell
	var err error

	switch len(dateString) {
	case 0:
		events, err = c.yandexoidUsecase.GetEvents(login)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		break

	default:
		date, err := time.Parse(time.DateOnly, dateString)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		events, err = c.yandexoidUsecase.GetEventsByDate(login, date)
	}

	return ctx.JSON(events)
}

func (c *YandexoidController) GetPickInfo(ctx *fiber.Ctx) error {
	key := ctx.Query("key")
	if len(key) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "'key' is not set")
	}

	login, err := c.giftUsecase.Get(key)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	yandexoid, err := c.yandexoidUsecase.Get(login.LuckyLogin)
	if err != nil {
		yandexoid = &domain.Yandexoid{
			Login:   login.LuckyLogin,
			Name:    "",
			Surname: "",
		}

		err := c.yandexoidUsecase.Create(yandexoid)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	events, err := c.yandexoidUsecase.GetEventsByDate(login.LuckyLogin, time.Now())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	pick := &domain.CreatePickReq{
		YandexoidLogin:    yandexoid.Login,
		EventUuid:         uuid.Nil,
		WithFriends:       false,
		IsListSuccess:     false,
		IsGiftSuccess:     false,
		IsBraceletSuccess: false,
		PickedAt:          time.Now(),
	}
	id, err := c.pickUsecase.CreatePick(pick)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	yandexoidRegs := &domain.YandexoidRegs{
		Yandexoid: yandexoid,
		Key:       key,
		PickId:    id,
		Events:    events,
	}

	return ctx.JSON(yandexoidRegs)
}
