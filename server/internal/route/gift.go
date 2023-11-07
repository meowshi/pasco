package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/meowshi/pasco-server/internal/controller"
	"github.com/meowshi/pasco-server/internal/storage/postgres"
	"github.com/meowshi/pasco-server/internal/usecase"
)

func SetupGiftRouter(app *fiber.App, db *sqlx.DB) {
	ps := postgres.NewPickPostgresStorage(db)

	gu := usecase.NewGiftUsecase()
	pu := usecase.NewPickUsecase(ps)
	c := controller.NewGiftController(gu, pu)

	router := app.Group("/gift")

	router.Get("", c.Get)
	router.Post("/:login", c.Give)
}
