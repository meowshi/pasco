package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/meowshi/pasco-server/internal/controller"
	"github.com/meowshi/pasco-server/internal/storage/postgres"
	"github.com/meowshi/pasco-server/internal/usecase"
	"google.golang.org/api/sheets/v4"
)

func SetupEnvRouter(app *fiber.App, db *sqlx.DB, service *sheets.Service) {
	s := postgres.NewEnvPostgresStorage(db)
	u := usecase.NewEnvUsecase(s, service)
	c := controller.NewEnvController(u)

	router := app.Group("/env")

	router.Get("", c.Get)
	router.Patch("/title", c.UpdateSheetTitle)
}
