package main

import (
  "github.com/gofiber/fiber"
  "github.com/QSOLink/Qsolink/qso"
)


func setupRoutes(app *fiber.App){
  app.Get("/api/v1/qso", getQsos)
  app.Get("/api/v1/qso/:id", getQso)
  app.Post("/api/v1/qso", postQso)
  app.Put("/api/v1/qso/:id", putQso)
  app.Delete("/api/v1/qso/:id", deleteQso)
}

func main() {
  app := fiber.New()

  setupRoutes(app)
  app.Listen(3000)
}


