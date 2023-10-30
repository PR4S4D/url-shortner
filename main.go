package main

import (
	"log"
	"url-shortner/controllers"
	"url-shortner/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	database.InitializeDB()

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")

	app.Get("/url", controllers.GetAllUrlData)
	app.Post("/shorten", controllers.Shorten)
	app.Get(":short_url", controllers.GetUrlData)

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", nil)
	})

	log.Fatal(app.Listen("127.0.0.1:9000"))

}
