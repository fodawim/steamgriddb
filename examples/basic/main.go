package main

import (
	"fmt"
	"sgdb"
)

func main() {

	app := sgdb.App{
		ApiKey: "PUT_API_KEY_HERE",
	}
	client, err := sgdb.NewClient(&app)
	if err != nil {
		panic(err)
	}

	game, err := client.Games.BySteamID("1794680")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Game: %s", game.Name)

}
