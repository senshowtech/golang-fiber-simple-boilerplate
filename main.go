package main

import (
	config "goFiber/config"
	dog "goFiber/controller/dog"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.Connect()
	app.Get("/", dog.GetDogs)
	app.Listen(":3000")
}
