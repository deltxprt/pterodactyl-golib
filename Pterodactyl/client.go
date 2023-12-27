package Pterodactyl

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"regexp"
	"strings"
)

type Client struct {
	URL    string `json:"url"`
	ApiKey string `json:"apikey"`
}

func (c *Client) ApiCall(path string, method string, body []byte, result interface{}) error {
	url := c.URL + path
	slog.Debug("url is: " + url)
	slog.Debug("method is: " + method)
	if body != nil {
		slog.Debug("body is: " + string(body))
	}

	request, err := http.NewRequest(method, url, strings.NewReader(string(body)))
	if err != nil {
		return fmt.Errorf("failed to create request: %s", err.Error())
	}
	slog.Debug("request is: ", request)

	request.Header.Set("Authorization", "Bearer "+c.ApiKey)
	request.Header.Set("Accept", "application/json")

	//fmt.Println("token is: " + c.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("failed to execute request: %s", err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error(err.Error(), "failed to close body")
		}
	}(resp.Body)

	slog.Debug("response: ",resp.Body)

	slog.Debug("response status: %d", resp.Status)

	err = json.NewDecoder(resp.Body).Decode(&result)

	if resp.StatusCode < 200 && resp.StatusCode > 299 {
		return fmt.Errorf("api call returned status code %d", resp.StatusCode)
	}
	if err != nil {
		return fmt.Errorf("failed to decode json response: %s", err.Error())
	}
	return nil
}

func NewClient(apiUrl string, token string) (client *Client, err error) {
	urlPattern := `^((http|https):\/\/)[\w\.\-]*$`
	match, err := regexp.MatchString(urlPattern, apiUrl)
	if err != nil {
		return nil, fmt.Errorf("error with regex pattern (%s)", urlPattern)
	}
	if !match {
		return nil, fmt.Errorf("invalid url: %s", apiUrl)
	}
	client = &Client{URL: apiUrl + "/api", ApiKey: token}

	return client, nil
}
