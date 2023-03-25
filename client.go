package sgdb

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	// UserAgent is the user agent used for all requests.
	UserAgent = "sgdb-go/0.1"

	// BaseURL is the base URL for the API.
	BaseURL = "https://www.steamgriddb.com/api/v2/"
)

// App holds necessary values to connect to the SteamGridDB api.
type App struct {
	// ApiKey is the SteamGridDB api key, which can be obtained from https://www.steamgriddb.com/profile/preferences/api
	ApiKey string
}

// Client manages access to the SteamGridDB api.
type Client struct {
	// Client is the HTTP client used to communicate with the SteamGridDB api.
	Client *http.Client

	// settings holds user-definable config options.
	settings *App

	// baseURL Base URL for the SteamGridDB api.
	baseURL *url.URL

	// Services used for talking to different parts of the SteamGridDB api.
	Games GameService
	Grids GridService
}

// NewClient returns a new SteamGridDB api client.
func NewClient(app *App) (*Client, error) {
	// Parse Base URL
	baseURL, err := url.Parse(BaseURL)
	if err != nil {
		return nil, err
	}

	// Setup the client
	client := &Client{
		Client:   http.DefaultClient,
		settings: app,
		baseURL:  baseURL,
	}

	// Setup services
	client.Games = &GameServiceOp{client: client}
	client.Grids = &GridServiceOp{client: client}

	return client, nil
}

// Get sends a GET request to the API and returns the response.
func (c *Client) Get(path string, res, opts interface{}) error {
	return c.Execute(http.MethodGet, path, nil, opts, res)
}

// Execute creates a request and sends it to the API, then returns the data or errors.
func (c *Client) Execute(method, path string, body, opts, res interface{}) error {
	request, err := c.NewRequest(method, path, body, opts)
	if err != nil {
		return err
	}
	return c.Do(request, res)
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, in which case it is resolved relative
// to the BaseURL of the Client.
func (c *Client) NewRequest(method, path string, body, opts interface{}) (*http.Request, error) {
	// Parse the URL
	providedURL, err := c.baseURL.Parse(path)
	if err != nil {
		return nil, err
	}
	u := c.baseURL.ResolveReference(providedURL)

	if opts != nil {

	}

	// Prepare the request body
	var bodyData []byte
	if body != nil {
		bodyData, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	// Create the request
	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, err
	}

	// Set the headers
	req.Header.Set("Authorization", "Bearer "+c.settings.ApiKey)
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

// Do sends a request to the API and returns the response.
func (c *Client) Do(req *http.Request, v interface{}) error {
	// Send the request
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Bad Request error handling (400 doesn't have a response body)
	if resp.StatusCode == http.StatusBadRequest {
		return err
	}

	// Data based error handling
	baseError := ResponseError{}
	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(bodyData) > 0 {
		err = json.Unmarshal(bodyData, &baseError)
		if err != nil {
			return err
		}
		if !baseError.Success {
			return &baseError
		}
	}

	// Decode the response
	err = json.NewDecoder(bytes.NewReader(bodyData)).Decode(v)
	if err != nil {
		return err
	}

	return err
}
