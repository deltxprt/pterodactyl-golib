package Pterodactyl

import (
	"encoding/json"
	"fmt"
)

type Allocations struct {
	Object string `json:"object"`
	Data   []struct {
		Object     string `json:"object"`
		Attributes struct {
			Id       int    `json:"id"`
			Ip       string `json:"ip"`
			Alias    string `json:"alias"`
			Port     int    `json:"port"`
			Notes    string `json:"notes"`
			Assigned bool   `json:"assigned"`
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
				Next string `json:"next"`
			} `json:"links"`
		} `json:"pagination"`
	} `json:"meta"`
}

type CreateAllocationsRequest struct {
	Ip    string   `json:"ip"`
	Ports []string `json:"ports"`
}

const allocationsPath = "/application/nodes/%d/allocations"

func (c *Client) GetAllocations(nodeId int) (Allocations, error) {
	var allocations Allocations
	err := c.ApiCall(fmt.Sprintf(allocationsPath, nodeId), "GET", nil, &allocations)
	if err != nil {
		return allocations, err
	}
	return allocations, nil
}

func (c *Client) CreateAllocations(nodeId int, request CreateAllocationsRequest) (Allocations, error) {
	var allocations Allocations
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return allocations, err
	}
	err = c.ApiCall(fmt.Sprintf(allocationsPath, nodeId), "POST", jsonRequest, &allocations)
	if err != nil {
		return allocations, err
	}
	return allocations, nil
}

func (c *Client) DeleteAllocation(nodeId int, allocationId int) error {
	err := c.ApiCall(fmt.Sprintf("%s/%d", fmt.Sprintf(allocationsPath, nodeId), allocationId), "DELETE", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
