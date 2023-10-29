package main

import (
	"log"
	"url-shortner/controllers"
	"url-shortner/database"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	database.InitializeDB()

	app := fiber.New()
	v1 := app.Group("api/v1")
	v1.Get("/url", controllers.GetAllUrlData)
	v1.Post("/shorten", controllers.Shorten)
	v1.Get(":short_url", controllers.GetUrlData)

	log.Fatal(app.Listen(":9000"))

}
