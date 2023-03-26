package sgdb

import (
	"fmt"
	"strings"
)

type GridService interface {
	// ByID returns a grid by its SteamGridDB ID.
	ByID(id string, opts interface{}) ([]Image, error)

	// ByPlatformID returns a grid by its platform ID.
	ByPlatformID(platform platform, ids []string, params interface{}) ([]Image, error)
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
	Items []Image `json:"data"`
}

type GridServiceOp struct {
	client *Client
}

func (s *GridServiceOp) ByID(id string, params interface{}) ([]Image, error) {
	path := fmt.Sprintf("grids/game/%s", id)
	grid := new(GridResponse)
	err := s.client.Get(path, grid, params)
	return grid.Items, err
}

func (s *GridServiceOp) ByPlatformID(platform platform, ids []string, params interface{}) ([]Image, error) {
	path := fmt.Sprintf("grids/%s/%s", platform, strings.Join(ids, ","))
	grid := new(GridResponse)
	err := s.client.Get(path, grid, params)
	return grid.Items, err
}
