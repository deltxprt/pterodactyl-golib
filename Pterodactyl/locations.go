package Pterodactyl

import (
	"encoding/json"
	"fmt"
	"time"
)

type Locations struct {
	Object string `json:"object"`
	Data   []struct {
		Object     string `json:"object"`
		Attributes struct {
			Id        int         `json:"id"`
			Short     string      `json:"short"`
			Long      interface{} `json:"long"`
			UpdatedAt time.Time   `json:"updated_at"`
			CreatedAt time.Time   `json:"created_at"`
		} `json:"attributes"`
	} `json:"data"`
	Meta struct {
		Pagination struct {
			Total       int `json:"total"`
			Count       int `json:"count"`
			PerPage     int `json:"per_page"`
			CurrentPage int `json:"current_page"`
			TotalPages  int `json:"total_pages"`
			Links       struct {
			} `json:"links"`
		} `json:"pagination"`
	} `json:"meta"`
}

type Location struct {
	Object     string `json:"object"`
	Attributes struct {
		Id        int         `json:"id"`
		Short     string      `json:"short"`
		Long      interface{} `json:"long"`
		UpdatedAt time.Time   `json:"updated_at"`
		CreatedAt time.Time   `json:"created_at"`
	} `json:"attributes"`
	Meta struct {
		Resource string `json:"resource,omitempty"`
	} `json:"meta,omitempty"`
}

const locationsPath = "/application/locations"

func GetLocations(pterodactylCfg PterodactylConfig) (Locations, error) {
	var result Locations
	err := ApiCall(pterodactylCfg, locationsPath, "GET", nil, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetLocation(pterodactylCfg PterodactylConfig, id int) (Location, error) {
	var result Location
	err := ApiCall(pterodactylCfg, fmt.Sprintf("%s/%d", locationsPath, id), "GET", nil, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func CreateLocation(pterodactylCfg PterodactylConfig, short string, long string) (Location, error) {
	var result Location
	var body = map[string]string{
		"short": short,
		"long":  long,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return result, err
	}
	err = ApiCall(pterodactylCfg, locationsPath, "POST", jsonBody, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func UpdateLocation(pterodactylCfg PterodactylConfig, id int, short string, long string) (Location, error) {
	var result Location
	var body = map[string]string{
		"short": short,
		"long":  long,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return result, err
	}
	err = ApiCall(pterodactylCfg, fmt.Sprintf("%s/%d", locationsPath, id), "PATCH", jsonBody, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func DeleteLocation(pterodactylCfg PterodactylConfig, id int) error {
	err := ApiCall(pterodactylCfg, fmt.Sprintf("%s/%d", locationsPath, id), "DELETE", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
