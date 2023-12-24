package Pterodactyl

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

type PterodactylConfig struct {
	URL    string `json:"url"`
	ApiKey string `json:"apikey"`
}

func ApiCall(pterodactylCfg PterodactylConfig, path string, method string, body []byte, result interface{}) error {
	url := pterodactylCfg.URL + path

	request, err := http.NewRequest(method, url, strings.NewReader(string(body)))
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", pterodactylCfg.ApiKey)

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
