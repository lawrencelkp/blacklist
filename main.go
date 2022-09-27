package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

var blacklist = []string{
	"5123fa9b3f0a8c1c193ee33d4fb0244a31e6cdf36a6421e78d317485e9aeb699",
	"b4e4d8f6a355d01617028dcbe210dae5460544def50ba463c69c84a0c60ac638",
	"0ba9fa9c3568f24ef6420952f322949f62e518b97cbfaf94bd9eaf9cb0b1a362",
	"fe0cdb29a83910b3e8f2cab3866b7f8e4867c6d8c2bfa6e1b16c0059199eaced",
	"717c0afe612ebea8b29c12dcbac5ffa60678dbaa42594f370a347ed3a8627b3b",
	"bda348ad58796fe53ab05acfa004a325d899f63e688983da3c8e88aa844588e7",
	"9a903cea361221848ce60ed15e376fb7362f54ed311d31f53183c638741bac8b",
	"02892bf3115b0c112832b0bf745b4b415a8eca79c8657781fdec31832f1d0fe0",
	"70d6607932a009b86190ae1a458e6e21a8d4ab1cb363a56efb5a058fc1b0fdd3",
	"d9788910a6b2a471f1d85fd4bbcd0c44b54bc53623ba31e955f08d48bb0f7c71",
}

type BlacklistRequest struct {
	AppId string `json:"appId"`
	Hash  string `json:"hash"`
}

type BlacklistResponse struct {
	Exists bool `json:"exists"`
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	app := fiber.New()

	// POST /api/blacklist
	app.Post("/api/blacklist", func(c *fiber.Ctx) error {
		blacklistRequest := new(BlacklistRequest)
		if err := c.BodyParser(blacklistRequest); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		hash := blacklistRequest.Hash

		if contains(blacklist, hash) {
			return c.Status(200).JSON(BlacklistResponse{true})
		} else {
			return c.Status(200).JSON(BlacklistResponse{false})
		}
	})

	log.Fatal(app.Listen(":3000"))
}
