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
			Id       int         `json:"id"`
			Ip       string      `json:"ip"`
			Alias    interface{} `json:"alias"`
			Port     int         `json:"port"`
			Notes    interface{} `json:"notes"`
			Assigned bool        `json:"assigned"`
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

func GetAllocations(pterodactylCfg PterodactylConfig, nodeId int) (Allocations, error) {
	var allocations Allocations
	err := ApiCall(pterodactylCfg, fmt.Sprintf(allocationsPath, nodeId), "GET", nil, &allocations)
	if err != nil {
		return allocations, err
	}
	return allocations, nil
}

func CreateAllocations(pterodactylCfg PterodactylConfig, nodeId int, request CreateAllocationsRequest) (Allocations, error) {
	var allocations Allocations
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return allocations, err
	}
	err = ApiCall(pterodactylCfg, fmt.Sprintf(allocationsPath, nodeId), "POST", jsonRequest, &allocations)
	if err != nil {
		return allocations, err
	}
	return allocations, nil
}

func DeleteAllocation(pterodactylCfg PterodactylConfig, nodeId int, allocationId int) error {
	err := ApiCall(pterodactylCfg, fmt.Sprintf("%s/%d", fmt.Sprintf(allocationsPath, nodeId), allocationId), "DELETE", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
