package sgdb

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGridServiceByID(t *testing.T) {
	setupTests()
	defer teardownTests()

	httpmock.RegisterResponder(
		http.MethodGet,
		"https://www.steamgriddb.com/api/v2/grids/game/123456789",
		httpmock.NewBytesResponder(
			http.StatusOK,
			loadTest("grid_by_id.json"),
		),
	)

	grids, err := client.Grids.ByID("123456789", nil)
	expectedImage := Grid{
		ID:        123456789,
		Score:     0,
		Style:     StyleAlternate,
		Width:     600,
		Height:    900,
		Nsfw:      false,
		Humor:     false,
		Mime:      MimePNG,
		Language:  "en",
		URL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/a.png",
		Thumb:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/thumb/a.jpg",
		Lock:      false,
		Epilepsy:  false,
		Upvotes:   0,
		Downvotes: 0,
		Author: &Author{
			Name:    "Author Name",
			Steam64: "1234512345",
			Avatar:  "https://avatars.akamai.steamstatic.com/4_medium.jpg",
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(grids))
		assert.Equal(t, expectedImage, grids[0])
	}

}

func TestGridServiceByPlatformID(t *testing.T) {
	setupTests()
	defer teardownTests()

	httpmock.RegisterResponder(
		http.MethodGet,
		"https://www.steamgriddb.com/api/v2/grids/steam/123456789",
		httpmock.NewBytesResponder(
			http.StatusOK,
			loadTest("grid_by_id.json"),
		),
	)

	grids, err := client.Grids.ByPlatformID(PlatformSteam, []string{"123456789"}, nil)

	expectedImage := Grid{
		ID:        123456789,
		Score:     0,
		Style:     StyleAlternate,
		Width:     600,
		Height:    900,
		Nsfw:      false,
		Humor:     false,
		Mime:      MimePNG,
		Language:  "en",
		URL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/grid/a.png",
		Thumb:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/thumb/a.jpg",
		Lock:      false,
		Epilepsy:  false,
		Upvotes:   0,
		Downvotes: 0,
		Author: &Author{
			Name:    "Author Name",
			Steam64: "1234512345",
			Avatar:  "https://avatars.akamai.steamstatic.com/4_medium.jpg",
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(grids))
		assert.Equal(t, expectedImage, grids[0])
	}

}
