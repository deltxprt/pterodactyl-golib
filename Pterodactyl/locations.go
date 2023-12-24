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
}

type CreateLocation struct {
	Object     string `json:"object"`
	Attributes struct {
		Id        int       `json:"id"`
		Short     string    `json:"short"`
		Long      string    `json:"long"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"attributes"`
	Meta struct {
		Resource string `json:"resource"`
	} `json:"meta"`
}

const locationsPath = "/application/locations"

func (app *PterodactylConfig) GetLocations() (Locations, error) {
	var result Locations
	err := app.apiCall(locationsPath, "GET", nil, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (app *PterodactylConfig) GetLocation(id int) (Location, error) {
	var result Location
	err := app.apiCall(fmt.Sprintf("%s/%d", locationsPath, id), "GET", nil, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (app *PterodactylConfig) CreateLocation(short string, long string) (Location, error) {
	var result Location
	var body = map[string]string{
		"short": short,
		"long":  long,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return result, err
	}
	err = app.apiCall(locationsPath, "POST", jsonBody, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (app *PterodactylConfig) UpdateLocation(id int, short string, long string) (Location, error) {
	var result Location
	var body = map[string]string{
		"short": short,
		"long":  long,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return result, err
	}
	err = app.apiCall(fmt.Sprintf("%s/%d", locationsPath, id), "PATCH", jsonBody, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (app *PterodactylConfig) DeleteLocation(id int) error {
	err := app.apiCall(fmt.Sprintf("%s/%d", locationsPath, id), "DELETE", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
