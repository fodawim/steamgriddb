package sgdb

import (
	"fmt"
)

type SearchService interface {
	// Search returns a list of games matching the given search query.
	Search(query string) ([]Game, error)
}

type SearchServiceOp struct {
	client *Client
}

type SearchResponse struct {
	Items []Game `json:"data"`
}

// Search returns a list of games matching the given search query.
func (s *SearchServiceOp) Search(query string) ([]Game, error) {
	path := fmt.Sprintf("search/autocomplete/%s", query)
	games := new(SearchResponse)
	err := s.client.Get(path, games, nil)
	return games.Items, err
}
