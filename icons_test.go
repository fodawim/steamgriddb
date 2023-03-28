package sgdb

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestIconServiceByID(t *testing.T) {
	setupTests()
	defer teardownTests()

	httpmock.RegisterResponder(
		http.MethodGet,
		"https://www.steamgriddb.com/api/v2/icons/game/123456789",
		httpmock.NewBytesResponder(
			http.StatusOK,
			loadTest("icons_by_id.json"),
		),
	)

	icons, err := client.Icons.ByID("123456789", nil)
	expectedIcon := Icon{
		ID:        123456789,
		Score:     0,
		Style:     StylesIconOfficial,
		Width:     0,
		Height:    0,
		Nsfw:      false,
		Humor:     false,
		Mime:      MimeIconVND,
		Language:  "en",
		Notes:     ":)",
		URL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/3.ico",
		Thumb:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/3/32/512x512.png",
		Lock:      false,
		Epilepsy:  false,
		Upvotes:   0,
		Downvotes: 0,
		Author: &Author{
			Name:    "Test Author",
			Steam64: "1234512345",
			Avatar:  "https://avatars.akamai.steamstatic.com/C.jpg",
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(icons))
		assert.Equal(t, expectedIcon, icons[0])
	}

}

func TestIconServiceByPlatformID(t *testing.T) {
	setupTests()
	defer teardownTests()

	httpmock.RegisterResponder(
		http.MethodGet,
		"https://www.steamgriddb.com/api/v2/icons/steam/123456789",
		httpmock.NewBytesResponder(
			http.StatusOK,
			loadTest("icons_by_id.json"),
		),
	)

	icons, err := client.Icons.ByPlatformID(PlatformSteam, []string{"123456789"}, nil)
	expectedIcon := Icon{
		ID:        123456789,
		Score:     0,
		Style:     StylesIconOfficial,
		Width:     0,
		Height:    0,
		Nsfw:      false,
		Humor:     false,
		Mime:      MimeIconVND,
		Notes:     ":)",
		Language:  "en",
		URL:       "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/3.ico",
		Thumb:     "https://cdn2.steamgriddb.com/file/sgdb-cdn/icon/3/32/512x512.png",
		Lock:      false,
		Epilepsy:  false,
		Upvotes:   0,
		Downvotes: 0,
		Author: &Author{
			Name:    "Test Author",
			Steam64: "1234512345",
			Avatar:  "https://avatars.akamai.steamstatic.com/C.jpg",
		},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(icons))
		assert.Equal(t, expectedIcon, icons[0])
	}

}
