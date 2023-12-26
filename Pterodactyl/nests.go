package Pterodactyl

import (
	"fmt"
	"time"
)

type Nests struct {
	Object string `json:"object"`
	Data   []struct {
		Object     string `json:"object"`
		Attributes struct {
			Id          int       `json:"id"`
			Uuid        string    `json:"uuid"`
			Author      string    `json:"author"`
			Name        string    `json:"name"`
			Description *string   `json:"description"`
			CreatedAt   time.Time `json:"created_at"`
			UpdatedAt   time.Time `json:"updated_at"`
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

type Nest struct {
	Object     string `json:"object"`
	Attributes struct {
		Id          int       `json:"id"`
		Uuid        string    `json:"uuid"`
		Author      string    `json:"author"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"attributes"`
}

const nestPath = "/application/nests"

func (c *Client)  GetNests() (Nests, error) {
	var nests Nests
	err := c.ApiCall(nestPath, "GET", nil, &nests)
	if err != nil {
		return nests, err
	}
	return nests, nil
}

func (c *Client) GetNest(id int) (Nest, error) {
	var nest Nest
	url := fmt.Sprintf("%s/%d", nestPath, id)
	err := c.ApiCall(url, "GET", nil, &nest)
	if err != nil {
		return nest, err
	}
	return nest, nil
}
