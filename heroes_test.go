package sgdb

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHeroServiceByID(t *testing.T) {
	setupTests()
	defer teardownTests()

	httpmock.RegisterResponder(
		http.MethodGet,
		"https://www.steamgriddb.com/api/v2/heroes/game/123456789",
		httpmock.NewBytesResponder(
			http.StatusOK,
			loadTest("heroes_by_id.json"),
		),
	)

	heroes, err := client.Heroes.ByID("123456789", nil)
	expectedImage := Hero{
		ID:        123456789,
		Score:     0,
		Style:     StylesHeroAlternate,
		Width:     1920,
		Height:    620,
		Nsfw:      false,
		Humor:     false,
		Mime:      MimeHeroPNG,
		Language:  "en",
		URL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/b.png",
		Thumb:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero_thumb/b.jpg",
		Lock:      false,
		Epilepsy:  false,
		Upvotes:   0,
		Downvotes: 0,
		Author: &Author{
			Name:    "Test Author",
			Steam64: "1234512345",
			Avatar:  "https://avatars.akamai.steamstatic.com/2.jpg",
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(heroes))
		assert.Equal(t, expectedImage, heroes[0])
	}

}

func TestHeroServiceByPlatformID(t *testing.T) {
	setupTests()
	defer teardownTests()

	httpmock.RegisterResponder(
		http.MethodGet,
		"https://www.steamgriddb.com/api/v2/heroes/steam/123456789",
		httpmock.NewBytesResponder(
			http.StatusOK,
			loadTest("heroes_by_id.json"),
		),
	)

	heroes, err := client.Heroes.ByPlatformID(PlatformSteam, []string{"123456789"}, nil)
	expectedImage := Hero{
		ID:        123456789,
		Score:     0,
		Style:     StylesHeroAlternate,
		Width:     1920,
		Height:    620,
		Nsfw:      false,
		Humor:     false,
		Mime:      MimeHeroPNG,
		Language:  "en",
		URL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero/b.png",
		Thumb:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/hero_thumb/b.jpg",
		Lock:      false,
		Epilepsy:  false,
		Upvotes:   0,
		Downvotes: 0,
		Author: &Author{
			Name:    "Test Author",
			Steam64: "1234512345",
			Avatar:  "https://avatars.akamai.steamstatic.com/2.jpg",
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(heroes))
		assert.Equal(t, expectedImage, heroes[0])
	}

}
