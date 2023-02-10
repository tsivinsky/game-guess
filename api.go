package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GameScreenshot struct {
	Id    uint   `json:"id"`
	Image string `json:"image"`
}

type GameScreenshotsResponse struct {
	Count    uint             `json:"count"`
	Next     *string          `json:"next"`
	Previous *string          `json:"previous"`
	Results  []GameScreenshot `json:"results"`
}

func GetGameScreenshots(id uint) (*GameScreenshotsResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.rawg.io/api/games/%d/screenshots?key=%s", id, Env.RAWG_API_KEY))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gameScreenshotsResp := new(GameScreenshotsResponse)
	if err = json.NewDecoder(resp.Body).Decode(&gameScreenshotsResp); err != nil {
		return nil, err
	}

	return gameScreenshotsResp, nil
}

type Game struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func GetGameById(id uint) (*Game, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.rawg.io/api/games/%d?key=%s", id, Env.RAWG_API_KEY))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	game := new(Game)
	if err = json.NewDecoder(resp.Body).Decode(&game); err != nil {
		return nil, err
	}

	return game, nil
}
