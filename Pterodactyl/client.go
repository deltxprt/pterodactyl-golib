package Pterodactyl

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

type Client struct {
	URL    string `json:"url"`
	ApiKey string `json:"apikey"`
}

func (c *Client) ApiCall(path string, method string, body []byte, result interface{}) error {
	url := c.URL + path

	request, err := http.NewRequest(method, url, strings.NewReader(string(body)))
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", c.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error(err.Error(), "failed to close body")
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		return fmt.Errorf("api call returned status code %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return err
	}
	return nil
}

func NewClient(apiUrl string, token string) (client *Client, err error) {

	client = &Client{URL: apiUrl, ApiKey: token}

	return client, nil
}