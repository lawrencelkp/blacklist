package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

var blacklist = []string{
	"0bf03569ba448bce827665cfe22e71d0357e0b3a2a17e9a86c3842c7596c875b",
	"86c80bf27ddc1f41b7758453112671dade338395a0a20c51f4c809a859422a19",
	"e83102429390a190f0192400649b2502965bf1dff7c625e2b3e69e698d2bf88f",
	"482f3e38e2eebd442b519c2ceb674922bbaaa33eeba6317cac18f12c76d744b6",
	"b9bb3a51709cfa448807bb4c586b5e23bc4d0d81a8a5dedd28bf3ef458f5603a",
	"21684d67722d7b6b86d1fb2cd179c637be4ee6242f10f1470973bea12bbaf911",
	"c0d259b7371c99f2cbbde91f4d86c711a6029a12a864b3ba71e9cbd3125b945e",
	"1251394118e37e206efe00e304cb8fdd6ed229827daf0725bd2430fc37a0181b",
	"56519acdfa337b7b09ad616ca16e18ba3fe1f3c9cd12204a8d88f5dd6ad5664e",
	"4ad0f704879e6d55a4e2f882e78dfbe397047d5122f6fd6022a006354a3d2a59",
	"59e2b57c9b46b8759e7f592900dde06f68ff6d8898a0d2031f90631761430bd4",
	"ecace5cc46045df72e6a9daf0e9c694236a829d5c78b3ab69bb043ac4a00a55c",
	"e2ee6173071f4c08d4e33a925cc080ac5919a27179a7a71aab2a53ab644cd03a",
	"e39e7e569ba1ff499075752706a3855610e2cce7e7e244951e2e6afad67f9e92",
	"5a44253bdce81b8511245187887b8f729507bea34894bbdfe9113e7ae6add946",
	"7e02c208edda35ace330438b430e674e6f551dd39a7ecdc6af71e351393bb83c",
	"1b1671630501a60538277d992846c21fe1e7347dd7686b5f273b6c3527d3eedd",
	"654df5c1022673cee810552e4ed212a10dae255199776fd447e02d84c413abd5",
	"01ff3946e06bfc446082698e51a1a078ff1157799ef0e35dafd035c62fe5cc4e",
	"4c5b56db336b61e12d261da5663f564ebf633435d04fcb7ae6e7b83e29f6dce6",
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

	app.Static("/web", "dist")

	app.Get("/web/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./dist/index.html")
	})

	// POST /api/mecCheck
	app.Post("/api/mecCheck", func(c *fiber.Ctx) error {
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

	log.Fatal(app.Listen(":8080"))
}
