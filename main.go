package main

import (
	config "goFiber/config"
	auth "goFiber/controller/auth"
	dog "goFiber/controller/dog"
	middleware "goFiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.Connect()
	app.Post("/login", auth.Login)
	app.Get("/", dog.GetDogs)
	app.Get("/restricted", middleware.AuthRequired(), auth.Restricted)
	app.Listen(":3000")
}
