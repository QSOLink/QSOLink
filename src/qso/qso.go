package qso

import (
  "github.com/gofiber/fiber"
)

func getQsos(c *fiber.Ctx) {
  c.Send("All QSOs")
}

func getQso(c *fiber.Ctx) {
  c.Send("Single QSO")
}

func postQso(c *fiber.Ctx) {
  c.Send("New QSO")
}

func putQso(c *fiber.Ctx) {
  c.Send("Update QSO")
}

func deleteQso(c *fiber.Ctx) {
  c.Send("Delete QSO")
}
