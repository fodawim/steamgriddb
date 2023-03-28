package sgdb

import (
	"fmt"
	"strings"
)

type HeroService interface {
	// ByID returns a hero by its SteamGridDB ID.
	ByID(id string, opts interface{}) ([]Hero, error)
	// ByPlatformID returns a hero by its platform ID.
	ByPlatformID(platform platform, ids []string, params interface{}) ([]Hero, error)
}

type HeroOptions struct {
	// Page params
	Page int `url:"page,omitempty"`

	// Filter params
	Style      StylesHero     `url:"styles,omitempty"`
	Dimensions DimensionsHero `url:"dimensions,omitempty"`
	MimeTypes  MimesHero      `url:"mimes,omitempty"`
	Types      Types          `url:"types,omitempty"`
	NSFW       trilean        `url:"nsfw,omitempty"`
	Humorous   trilean        `url:"humor,omitempty"`
	Epileptic  trilean        `url:"epilepsy,omitempty"`
	Tags       Tags           `url:"oneoftag,omitempty"`
}

type HeroResponse struct {
	Items []Hero `json:"data"`
}

type Hero struct {
	ID        int       `json:"id"`
	Score     int       `json:"score"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Nsfw      bool      `json:"nsfw"`
	Humor     bool      `json:"humor"`
	Notes     string    `json:"notes"`
	Language  string    `json:"language"`
	URL       string    `json:"url"`
	Thumb     string    `json:"thumb"`
	Lock      bool      `json:"lock"`
	Epilepsy  bool      `json:"epilepsy"`
	Upvotes   int       `json:"upvotes"`
	Downvotes int       `json:"downvotes"`
	Author    *Author   `json:"author"`
	Style     herostyle `json:"style"`
	Mime      mimehero  `json:"mime"`
}

type HeroServiceOp struct {
	client *Client
}

func (s *HeroServiceOp) ByID(id string, params interface{}) ([]Hero, error) {
	path := fmt.Sprintf("heroes/game/%s", id)
	heroes := new(HeroResponse)
	err := s.client.Get(path, heroes, params)
	return heroes.Items, err
}

func (s *HeroServiceOp) ByPlatformID(platform platform, ids []string, params interface{}) ([]Hero, error) {
	path := fmt.Sprintf("heroes/%s/%s", platform, strings.Join(ids, ","))
	heroes := new(HeroResponse)
	err := s.client.Get(path, heroes, params)
	return heroes.Items, err
}
