package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/meowshi/pasco-server/pkg/domain"
	"github.com/meowshi/pasco-server/pkg/usecase"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type EventController struct {
	usecase usecase.EventUsecase
}

func NewEventController(usecase usecase.EventUsecase) *EventController {
	return &EventController{
		usecase: usecase,
	}
}

type EventCreateReq struct {
	GoogleSheetCell string `json:"google_sheet_cell"`
}

func (c *EventController) Create(ctx *fiber.Ctx) error {
	req := &EventCreateReq{}

	err := ctx.BodyParser(req)
	if err != nil {
		logrus.Errorf("Ошибка при парсинге тела: %s.", err)

		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = c.usecase.Create(req.GoogleSheetCell)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func (c *EventController) Delete(ctx *fiber.Ctx) error {
	uuid, err := uuid.Parse(ctx.Params("uuid"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = c.usecase.Delete(uuid)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(fasthttp.StatusOK)
}

func (c *EventController) Get(ctx *fiber.Ctx) error {
	uuid, err := uuid.Parse(ctx.Params("uuid"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	event, err := c.usecase.Get(uuid)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(event)
}

func (c *EventController) GetByDate(ctx *fiber.Ctx) error {
	dateString := ctx.Query("date")

	var date time.Time
	var err error
	switch len(dateString) {
	case 0:
		date = time.Now()
		break
	default:
		date, err = time.Parse(time.DateOnly, dateString)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	}

	events, err := c.usecase.GetByDate(date)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(events)
}

func (c *EventController) Update(ctx *fiber.Ctx) error {
	var event = &domain.Event{}

	err := ctx.BodyParser(event)
	if err != nil {
		logrus.Errorf("Ошибка при парсинге тела эвента: %s.", err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = c.usecase.Update(event)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c *EventController) GetEventLists(ctx *fiber.Ctx) error {
	uuidString := ctx.Params("uuid")
	uuid, err := uuid.Parse(uuidString)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	list, err := c.usecase.GetEventLists(uuid)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(list)
}
