package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/meowshi/pasco-server/internal/controller"
	"github.com/meowshi/pasco-server/internal/storage/postgres"
	"github.com/meowshi/pasco-server/internal/usecase"
)

func SetupYandexoidRouter(app *fiber.App, db *sqlx.DB) {
	yandexoidStorage := postgres.NewYandexoidPostgresStorage(db)
	pickStorage := postgres.NewPickPostgresStorage(db)
	yandexoidUsecase := usecase.NewYandexoidUsecase(yandexoidStorage)
	pickUsecase := usecase.NewPickUsecase(pickStorage)
	giftUsecase := usecase.NewGiftUsecase()
	yandexoidController := controller.NewYandexoidController(yandexoidUsecase, giftUsecase, pickUsecase)

	router := app.Group("/yandexoid")

	router.Get("/:login/events", yandexoidController.GetEvents)
	router.Get("", yandexoidController.GetPickInfo)
}
