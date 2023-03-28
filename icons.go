package sgdb

import (
	"fmt"
	"strings"
)

type IconService interface {
	// ByID returns a Icon by its SteamGridDB ID.
	ByID(id string, opts interface{}) ([]Icon, error)
	// ByPlatformID returns a Icon by its platform ID.
	ByPlatformID(platform platform, ids []string, params interface{}) ([]Icon, error)
}

type IconOptions struct {
	// Page params
	Page int `url:"page,omitempty"`

	// Filter params
	Style      StylesIcon     `url:"styles,omitempty"`
	Dimensions DimensionsIcon `url:"dimensions,omitempty"`
	MimeTypes  MimesLogo      `url:"mimes,omitempty"`
	Types      Types          `url:"types,omitempty"`
	NSFW       trilean        `url:"nsfw,omitempty"`
	Humorous   trilean        `url:"humor,omitempty"`
	Epileptic  trilean        `url:"epilepsy,omitempty"`
	Tags       Tags           `url:"oneoftag,omitempty"`
}

type IconResponse struct {
	Items []Icon `json:"data"`
}

type Icon struct {
	ID        int       `json:"id"`
	Score     int       `json:"score"`
	Style     iconstyle `json:"style"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Nsfw      bool      `json:"nsfw"`
	Humor     bool      `json:"humor"`
	Notes     string    `json:"notes"`
	Mime      mimeicon  `json:"mime"`
	Language  string    `json:"language"`
	URL       string    `json:"url"`
	Thumb     string    `json:"thumb"`
	Lock      bool      `json:"lock"`
	Epilepsy  bool      `json:"epilepsy"`
	Upvotes   int       `json:"upvotes"`
	Downvotes int       `json:"downvotes"`
	Author    *Author   `json:"author"`
}

type IconServiceOp struct {
	client *Client
}

func (s *IconServiceOp) ByID(id string, params interface{}) ([]Icon, error) {
	path := fmt.Sprintf("icons/game/%s", id)
	icons := new(IconResponse)
	err := s.client.Get(path, icons, params)
	return icons.Items, err
}

func (s *IconServiceOp) ByPlatformID(platform platform, ids []string, params interface{}) ([]Icon, error) {
	path := fmt.Sprintf("icons/%s/%s", platform, strings.Join(ids, ","))
	icons := new(IconResponse)
	err := s.client.Get(path, icons, params)
	return icons.Items, err
}
