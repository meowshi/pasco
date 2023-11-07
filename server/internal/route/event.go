package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/meowshi/pasco-server/internal/controller"
	"github.com/meowshi/pasco-server/internal/storage/postgres"
	internal_usecase "github.com/meowshi/pasco-server/internal/usecase"
	"github.com/meowshi/pasco-server/pkg/storage"
	"github.com/meowshi/pasco-server/pkg/usecase"
	"google.golang.org/api/sheets/v4"
)

func SetupEventRouter(app *fiber.App, db *sqlx.DB, service *sheets.Service) {
	var eventStorage storage.EventStorage = postgres.NewEventPostgresStorage(db)
	var envStorage storage.EnvStorage = postgres.NewEnvPostgresStorage(db)
	var yandexoidStorage storage.YandexoidStorage = postgres.NewYandexoidPostgresStorage(db)
	var registrationStorage storage.RegistrationStorage = postgres.NewRegistrationPostgresStorage(db)

	var usecase usecase.EventUsecase = internal_usecase.NewEventUsecase(service, eventStorage, registrationStorage, yandexoidStorage, envStorage)

	controller := controller.NewEventController(usecase)

	router := app.Group("/event")

	router.Post("", controller.Create)
	router.Get("", controller.GetByDate)
	router.Get("/:uuid", controller.Get)
	router.Get("/:uuid/lists", controller.GetEventLists)
	router.Delete("/:uuid", controller.Delete)
	router.Patch("", controller.Update)
}
