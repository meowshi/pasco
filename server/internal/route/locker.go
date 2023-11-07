package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/meowshi/pasco-server/internal/controller"
	"github.com/meowshi/pasco-server/internal/storage/postgres"
	"github.com/meowshi/pasco-server/internal/usecase"
)

func SetupLockerRouter(app *fiber.App, db *sqlx.DB) {
	ps := postgres.NewPickPostgresStorage(db)

	lu := usecase.NewLockerUsecase()
	pu := usecase.NewPickUsecase(ps)

	c := controller.NewLockerController(lu, pu)

	router := app.Group("/locker")

	router.Get("", c.GetLockerEvents)
	router.Post("/print", c.Print)
	router.Get("/printers", c.GetPrinters)
}
