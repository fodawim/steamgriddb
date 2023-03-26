package sgdb

import (
	"fmt"
	"strings"
)

type HeroService interface {
	// ByID returns a hero by its SteamGridDB ID.
	ByID(id string, opts interface{}) ([]Image, error)
	// ByPlatformID returns a hero by its platform ID.
	ByPlatformID(platform platform, ids []string, params interface{}) ([]Image, error)
}

type HeroOptions struct {
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

type HeroResponse struct {
	Items []Image `json:"data"`
}

type HeroServiceOp struct {
	client *Client
}

func (s *HeroServiceOp) ByID(id string, params interface{}) ([]Image, error) {
	path := fmt.Sprintf("heroes/game/%s", id)
	heroes := new(HeroResponse)
	err := s.client.Get(path, heroes, params)
	return heroes.Items, err
}

func (s *HeroServiceOp) ByPlatformID(platform platform, ids []string, params interface{}) ([]Image, error) {
	path := fmt.Sprintf("heroes/%s/%s", platform, strings.Join(ids, ","))
	heroes := new(HeroResponse)
	err := s.client.Get(path, heroes, params)
	return heroes.Items, err
}
