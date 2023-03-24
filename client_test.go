package sgdb

import (
	"github.com/jarcoal/httpmock"
	"os"
)

var (
	app    App
	client *Client
)

func setupTests() {
	// Fake the app, apikey not needed for testing
	app = App{
		ApiKey: "nope",
	}

	// Create a new client
	var err error
	client, err = NewClient(&app)
	if err != nil {
		panic(err)
	}

	// Mock the server responses
	httpmock.ActivateNonDefault(client.Client)
}

func teardownTests() {
	httpmock.Deactivate()
}

func loadTest(name string) []byte {
	// Load the test json file
	data, err := os.ReadFile("tests/" + name)
	if err != nil {
		panic(err)
	}

	return data
}
