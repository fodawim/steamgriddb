package sgdb

import (
	"fmt"
	"strings"
)

type LogoService interface {
	// ByID returns a Logo by its SteamGridDB ID.
	ByID(id string, opts interface{}) ([]Logo, error)
	// ByPlatformID returns a Logo by its platform ID.
	ByPlatformID(platform platform, ids []string, params interface{}) ([]Logo, error)
}

type LogoOptions struct {
	// Page params
	Page int `url:"page,omitempty"`

	// Filter params
	Style     StylesLogo `url:"styles,omitempty"`
	MimeTypes MimesLogo  `url:"mimes,omitempty"`
	Types     Types      `url:"types,omitempty"`
	NSFW      trilean    `url:"nsfw,omitempty"`
	Humorous  trilean    `url:"humor,omitempty"`
	Epileptic trilean    `url:"epilepsy,omitempty"`
	Tags      Tags       `url:"oneoftag,omitempty"`
}

type LogoResponse struct {
	Items []Logo `json:"data"`
}

type Logo struct {
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
	Style     logostyle `json:"style"`
	Mime      mimelogo  `json:"mime"`
}

type LogoServiceOp struct {
	client *Client
}

func (s *LogoServiceOp) ByID(id string, params interface{}) ([]Logo, error) {
	path := fmt.Sprintf("logos/game/%s", id)
	logos := new(LogoResponse)
	err := s.client.Get(path, logos, params)
	return logos.Items, err
}

func (s *LogoServiceOp) ByPlatformID(platform platform, ids []string, params interface{}) ([]Logo, error) {
	path := fmt.Sprintf("logos/%s/%s", platform, strings.Join(ids, ","))
	logos := new(LogoResponse)
	err := s.client.Get(path, logos, params)
	return logos.Items, err
}
