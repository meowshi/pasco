package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/meowshi/pasco-server/internal/route"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	databaseUrl := os.Getenv("DATABASE_URL")
	if len(databaseUrl) == 0 {
		panic("DATABASE_URL не установлена")
	}

	db, err := sqlx.Connect("pgx", databaseUrl)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(2)

	sheetService := SetupSheetService()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("index")
	})
	app.Use(func(c *fiber.Ctx) error {
		res := c.Response()
		res.Header.Add("Access-Control-Allow-Origin", "*")
		return c.Next()
	})
	app.Use(cors.New())

	route.SetupEventRouter(app, db, sheetService)
	route.SetupYandexoidRouter(app, db)
	route.SetupGiftRouter(app, db)
	route.SetupLockerRouter(app, db)
	route.SetupRegistrationRouter(app, db, sheetService)
	route.SetupPickRouter(app, db)
	route.SetupEnvRouter(app, db, sheetService)

	app.ListenTLS(":443", "cert.pem", "key.unencrypted.pem")
}

func SetupSheetService() *sheets.Service {
	ctx := context.Background()

	credentials, err := os.ReadFile("credentials.json")
	if err != nil {
		panic("Отсутсвует файл credentials.json")
	}

	config, err := google.JWTConfigFromJSON(credentials, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		panic("Что-то не так с credentials.json")
	}

	client := config.Client(ctx)

	service, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		panic(fmt.Sprintf("Unable to retrieve Sheets client: %v", err))
	}

	return service
}
