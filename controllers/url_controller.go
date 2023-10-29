package controllers

import (
	"url-shortner/database"
	"url-shortner/models"
	"url-shortner/repository"
	"url-shortner/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllUrlData(c *fiber.Ctx) error {
	urlData, err := repository.NewUrlRepository(database.DB).QueryAll()

	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(urlData)
}

func Shorten(c *fiber.Ctx) error {
	req := new(models.ShortenRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	shortUrl := utils.GenerateUniqueId(6)
	for {
		result, _ := repository.NewUrlRepository(database.DB).Query(shortUrl)
		if result.Id != -1 {
			shortUrl = utils.GenerateUniqueId(6)
		} else {
			break
		}
	}

	res, err := repository.NewUrlRepository(database.DB).Save(req.Url, shortUrl)
	response := new(models.Result)
	response.Success = res == 1
	response.Error = err
	response.ShortUrl = shortUrl

	return c.JSON(response)
}

func GetUrlData(c *fiber.Ctx) error {
	shortUrl := c.AllParams()["short_url"]
	result, _ := repository.NewUrlRepository(database.DB).Query(shortUrl)

	if result.Id == -1 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Redirect(result.Url)
}
