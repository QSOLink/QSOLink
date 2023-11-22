package main

import (
	"log"

	"github.com/QSOLink/QSOLink/database"
	_ "github.com/QSOLink/QSOLink/docs"
	"github.com/QSOLink/QSOLink/qso"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

//	@title			QSOLink
//	@description	JSON/REST based Ham Radio Logging Server
func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()
	defer database.DB.Close()

	api := app.Group("/api")
	qso.Register(api, database.DB)

	app.Get("/docs/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":5001"))
}
