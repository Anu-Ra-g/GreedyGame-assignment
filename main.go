package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func init() {
	InitializeKeystore()
	InitializeListStore()
}

func mux(c *fiber.Ctx) error {

	var action Command

	if err := c.BodyParser(&action); err != nil {
		fmt.Println(err)
		return err
	}

	list := strings.Split(action.Command, " ")

	go deleteExpiredKeys()

	switch list[0] {
	case "GET":
		value, err := get(list)
		_ = c.JSON(fiber.Map{
			"value": value,
		})
		if err != nil {
			_ = c.Status(404).JSON(fiber.Map{
				"error": "key doesnot exist",
			})
		}
	case "SET":
		set(list)
	case "QPUSH":
		qpush(list)
	case "QPOP":
		value, err := qpop(list)
		_ = c.JSON(fiber.Map{
			"value": value,
		})
		if err != nil {
			_ = c.Status(404).JSON(fiber.Map{
				"error": "key doesnot exist",
			})
		}
	default:
		return c.Status(400).JSON(fiber.Map{"message": "Invalid Command or Invalid Command"})
	}

	return nil
}

func main() {

	app := fiber.New()

	app.Get("/", mux)

	if err := app.Listen(":3000"); err != nil {
		fmt.Println(err)
	}
}
