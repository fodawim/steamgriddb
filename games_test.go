package sgdb

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

const (
	expectedGameID = 123456789
)

func TestGameServiceOpByID(t *testing.T) {
	setupTests()
	defer teardownTests()

	httpmock.RegisterResponder(
		http.MethodGet,
		"https://www.steamgriddb.com/api/v2/games/id/123456789",
		httpmock.NewBytesResponder(
			http.StatusOK,
			loadTest("games_by_id.json"),
		),
	)

	game, err := client.Games.ByID(strconv.Itoa(expectedGameID))

	expectedGame := &Game{
		ID:   expectedGameID,
		Name: "Counter-Strike: Global Offensive",
		Types: []string{
			"steam",
		},
		Verified: true,
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedGame, game)
}
