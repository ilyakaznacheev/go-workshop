package jokes

import (
	"fmt"
	"net/http"

	"gopkg.in/square/go-jose.v2/json"
)

const getJokePath = "/api?format=json"

// JokeClient is a joke API client
type JokeClient struct {
	url string
}

func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + getJokePath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request error: %v", err)
	}

	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return *data, nil
}
