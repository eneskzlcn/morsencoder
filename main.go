package main

import (
	"MorseCodeEncoder/encoder"
	"github.com/gofiber/fiber/v2"
)

func main(){
	app:= fiber.New()
	app.Get("/encode", func(c *fiber.Ctx) error {
		var sentTextToEncode string
		sentTextToEncode = c.Query("text")
		if sentTextToEncode != ""{
			return c.SendString(encoder.EncodeToMorseCode(sentTextToEncode," "))
		}
		return c.SendStatus(404)
	})
	app.Listen(":4006")
}
