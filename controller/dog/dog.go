package dog

import (
	config "goFiber/config"
	dog "goFiber/model/dog"

	"github.com/gofiber/fiber/v2"
)

func GetDogs(c *fiber.Ctx) error {
	var dogs []dog.Dog
	config.Database.Find(&dogs)
	return c.Status(200).JSON(dogs)
}
