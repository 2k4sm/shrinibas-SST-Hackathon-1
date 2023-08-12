package moviedb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	http   *http.Client
	apikey string
}

type Movies struct {
	Title      string `json:"title"`
	Year       string `json:"year"`
	Genre      string `json:"genre"`
	Director   string `json:"director"`
	Plot       string `json:"plot"`
	Poster     string `json:"poster"`
	Type       string `json:"type"`
	ImdbRating string `json:"imdbRating"`
}

func NewClient(httpClient *http.Client, key string) *Client {

	return &Client{httpClient, key}
}

func (c *Client) FetchMovie(query string) (*Movies, error) {
	endpoint := fmt.Sprintf("https://www.omdbapi.com/?apikey=%s&t=%s", c.apikey, url.QueryEscape(query))

	resp, err := c.http.Get(endpoint)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	res := &Movies{}

	return res, json.Unmarshal(body, res)

}
