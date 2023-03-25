package sgdb

import (
	"fmt"
	"strings"
)

type GridService interface {
	// ByID returns a grid by its SteamGridDB ID.
	ByID(id string, opts interface{}) ([]Grid, error)

	// ByPlatformID returns a grid by its platform ID.
	ByPlatformID(platform platform, ids []string, params interface{}) ([]Grid, error)
}

type GridOptions struct {
	// Page params
	Page int `url:"page,omitempty"`

	// Filter params
	Style      Styles     `url:"styles,omitempty"`
	Dimensions Dimensions `url:"dimensions,omitempty"`
	MimeTypes  Mimes      `url:"mimes,omitempty"`
	Types      Types      `url:"types,omitempty"`
	NSFW       trilean    `url:"nsfw,omitempty"`
	Humorous   trilean    `url:"humor,omitempty"`
	Epileptic  trilean    `url:"epilepsy,omitempty"`
	Tags       Tags       `url:"oneoftag,omitempty"`
}

type GridResponse struct {
	Items []Grid `json:"data"`
}

type Grid struct {
	ID        int     `json:"id"`
	Score     int     `json:"score"`
	Style     style   `json:"style"`
	Width     int     `json:"width"`
	Height    int     `json:"height"`
	Nsfw      bool    `json:"nsfw"`
	Humor     bool    `json:"humor"`
	Notes     string  `json:"notes"`
	Mime      mime    `json:"mime"`
	Language  string  `json:"language"`
	URL       string  `json:"url"`
	Thumb     string  `json:"thumb"`
	Lock      bool    `json:"lock"`
	Epilepsy  bool    `json:"epilepsy"`
	Upvotes   int     `json:"upvotes"`
	Downvotes int     `json:"downvotes"`
	Author    *Author `json:"author"`
}

type GridServiceOp struct {
	client *Client
}

func (s *GridServiceOp) ByID(id string, params interface{}) ([]Grid, error) {
	path := fmt.Sprintf("grids/game/%s", id)
	grid := new(GridResponse)
	err := s.client.Get(path, grid, params)
	return grid.Items, err
}

func (s *GridServiceOp) ByPlatformID(platform platform, ids []string, params interface{}) ([]Grid, error) {
	path := fmt.Sprintf("grids/%s/%s", platform, strings.Join(ids, ","))
	grid := new(GridResponse)
	err := s.client.Get(path, grid, params)
	return grid.Items, err
}
