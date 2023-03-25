package sgdb

import (
	"fmt"
)

// GameService allows access to the single game related functions in the SteamGridDB api.
type GameService interface {
	// ByID returns a game by its SteamGridDB ID.
	ByID(id string) (*Game, error)
	// BySteamID returns a game by its Steam App ID.
	BySteamID(id string) (*Game, error)
}

// Game represents a game on SteamGridDB.
type Game struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Types    []string `json:"types"`
	Verified bool     `json:"verified"`
}

// GameResponse represents a response from the SteamGridDB api.
type GameResponse struct {
	Game *Game `json:"data"`
}

// GameServiceOp handles communicating with to the game services from the SteamGridDB api.
type GameServiceOp struct {
	client *Client
}

// ByID returns a game using its SteamGridDB ID.
func (s *GameServiceOp) ByID(id string) (*Game, error) {
	path := fmt.Sprintf("games/id/%s", id)
	game := new(GameResponse)
	err := s.client.Get(path, game)
	return game.Game, err
}

// BySteamID returns a game using its Steam App ID.
func (s *GameServiceOp) BySteamID(id string) (*Game, error) {
	path := fmt.Sprintf("games/steam/%s", id)
	game := new(GameResponse)
	err := s.client.Get(path, game)
	return game.Game, err
}
