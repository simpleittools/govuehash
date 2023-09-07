package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword will hash the password using the bcrypt algorithm
func HashPassword(examplePassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(examplePassword), 12)
	if err != nil {
		fmt.Println(err)
	}
	return string(hashedPassword), nil
}

func main() {
	PORT := ":3030"
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	var examplePassword string
	var resultingHash string

	// POST route for posting the example ot the hashing algorithm
	app.Post("/api/hashing", func(c *fiber.Ctx) error {
		err := c.BodyParser(&examplePassword)
		collect, _ := json.Marshal(examplePassword)

		hashedPassword, err := HashPassword(string(collect))
		if err != nil {
			return err
		}
		resultingHash = hashedPassword

		//fmt.Printf("Your original password was " + examplePassword + "Your resulting password is " + resultingHash)
		return c.JSON(hashedPassword)

	})
	// GET route to display the result from the hashing algorithm
	app.Get("/api/result", func(c *fiber.Ctx) error {
		return c.SendString(resultingHash)
	})

	app.Listen(PORT)
}
