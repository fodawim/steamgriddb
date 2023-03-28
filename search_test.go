package sgdb

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSearchService(t *testing.T) {
	setupTests()
	defer teardownTests()

	httpmock.RegisterResponder(
		http.MethodGet,
		"https://www.steamgriddb.com/api/v2/search/autocomplete/test",
		httpmock.NewBytesResponder(
			http.StatusOK,
			loadTest("search.json"),
		),
	)

	games, err := client.Search.Search("test")
	expectedGame := Game{
		ID:          12345,
		Name:        "Test Game",
		Verified:    true,
		ReleaseDate: 1224547200,
		Types:       []string{},
	}

	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(games))
		assert.Equal(t, expectedGame, games[0])
	}

}
