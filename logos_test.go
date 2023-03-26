package sgdb

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestLogoServiceByID(t *testing.T) {
	setupTests()
	defer teardownTests()

	httpmock.RegisterResponder(
		http.MethodGet,
		"https://www.steamgriddb.com/api/v2/logos/game/123456789",
		httpmock.NewBytesResponder(
			http.StatusOK,
			loadTest("logos_by_id.json"),
		),
	)

	logos, err := client.Logos.ByID("123456789", nil)
	expectedLogo := Image{
		ID:        123456789,
		Score:     0,
		Style:     StyleOfficial,
		Width:     9000,
		Height:    1000,
		Nsfw:      false,
		Humor:     false,
		Mime:      MimePNG,
		Language:  "en",
		URL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/3.png",
		Thumb:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo_thumb/3.png",
		Lock:      false,
		Epilepsy:  false,
		Upvotes:   0,
		Downvotes: 0,
		Author: &Author{
			Name:    "Test Author",
			Steam64: "1234512345",
			Avatar:  "https://avatars.akamai.steamstatic.com/7.jpg",
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(logos))
		assert.Equal(t, expectedLogo, logos[0])
	}

}

func TestLogoServiceByPlatformID(t *testing.T) {
	setupTests()
	defer teardownTests()

	httpmock.RegisterResponder(
		http.MethodGet,
		"https://www.steamgriddb.com/api/v2/logos/steam/123456789",
		httpmock.NewBytesResponder(
			http.StatusOK,
			loadTest("logos_by_id.json"),
		),
	)

	logos, err := client.Logos.ByPlatformID(PlatformSteam, []string{"123456789"}, nil)
	expectedLogo := Image{
		ID:        123456789,
		Score:     0,
		Style:     StyleOfficial,
		Width:     9000,
		Height:    1000,
		Nsfw:      false,
		Humor:     false,
		Mime:      MimePNG,
		Language:  "en",
		URL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo/3.png",
		Thumb:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/logo_thumb/3.png",
		Lock:      false,
		Epilepsy:  false,
		Upvotes:   0,
		Downvotes: 0,
		Author: &Author{
			Name:    "Test Author",
			Steam64: "1234512345",
			Avatar:  "https://avatars.akamai.steamstatic.com/7.jpg",
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(logos))
		assert.Equal(t, expectedLogo, logos[0])
	}

}
