package main

import (
	"log"
	"math/rand"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/tsivinsky/goenv"
)

const MAX_ID = 868105

func main() {
	err := goenv.Load(Env)
	if err != nil {
		log.Fatal(err)
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		ids := GenerateRandomNumbers(3, MAX_ID)

		var games []*Game
		var wg sync.WaitGroup

		for _, id := range ids {
			wg.Add(1)

			go func(id int) {
				defer wg.Done()

				game, err := GetGameById(uint(id))
				if err != nil || game == nil {
					return
				}

				games = append(games, game)
			}(id)
		}

		wg.Wait()

		gameIdx := rand.Intn(len(games))
		game := games[gameIdx]

		gameScreenshots, err := GetGameScreenshots(game.Id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		i := rand.Intn(len(gameScreenshots.Results))
		randomScreenshot := gameScreenshots.Results[i]

		return c.Render("index", fiber.Map{
			"games":      games,
			"answer":     game.Id,
			"screenshot": randomScreenshot.Image,
		})
	})

	log.Fatal(app.Listen(":5000"))
}
