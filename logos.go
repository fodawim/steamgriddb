package sgdb

import (
	"fmt"
	"strings"
)

type LogoService interface {
	// ByID returns a Logo by its SteamGridDB ID.
	ByID(id string, opts interface{}) ([]Image, error)
	// ByPlatformID returns a Logo by its platform ID.
	ByPlatformID(platform platform, ids []string, params interface{}) ([]Image, error)
}

type LogoOptions struct {
	// Page params
	Page int `url:"page,omitempty"`

	// Filter params
	Style     Styles  `url:"styles,omitempty"`
	MimeTypes Mimes   `url:"mimes,omitempty"`
	Types     Types   `url:"types,omitempty"`
	NSFW      trilean `url:"nsfw,omitempty"`
	Humorous  trilean `url:"humor,omitempty"`
	Epileptic trilean `url:"epilepsy,omitempty"`
	Tags      Tags    `url:"oneoftag,omitempty"`
}

type LogoResponse struct {
	Items []Image `json:"data"`
}

type LogoServiceOp struct {
	client *Client
}

func (s *LogoServiceOp) ByID(id string, params interface{}) ([]Image, error) {
	path := fmt.Sprintf("logos/game/%s", id)
	logos := new(LogoResponse)
	err := s.client.Get(path, logos, params)
	return logos.Items, err
}

func (s *LogoServiceOp) ByPlatformID(platform platform, ids []string, params interface{}) ([]Image, error) {
	path := fmt.Sprintf("logos/%s/%s", platform, strings.Join(ids, ","))
	logos := new(LogoResponse)
	err := s.client.Get(path, logos, params)
	return logos.Items, err
}
