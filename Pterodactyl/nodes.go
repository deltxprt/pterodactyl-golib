package Pterodactyl

import (
	"encoding/json"
	"fmt"
	"time"
)

type Nodes struct {
	Object string `json:"object"`
	Data   []struct {
		Object     string `json:"object"`
		Attributes struct {
			Id                 int         `json:"id"`
			Uuid               string      `json:"uuid"`
			Public             bool        `json:"public"`
			Name               string      `json:"name"`
			Description        interface{} `json:"description"`
			LocationId         int         `json:"location_id"`
			Fqdn               string      `json:"fqdn"`
			Scheme             string      `json:"scheme"`
			BehindProxy        bool        `json:"behind_proxy"`
			MaintenanceMode    bool        `json:"maintenance_mode"`
			Memory             int         `json:"memory"`
			MemoryOverallocate int         `json:"memory_overallocate"`
			Disk               int         `json:"disk"`
			DiskOverallocate   int         `json:"disk_overallocate"`
			UploadSize         int         `json:"upload_size"`
			DaemonListen       int         `json:"daemon_listen"`
			DaemonSftp         int         `json:"daemon_sftp"`
			DaemonBase         string      `json:"daemon_base"`
			CreatedAt          time.Time   `json:"created_at"`
			UpdatedAt          time.Time   `json:"updated_at"`
			AllocatedResources struct {
				Memory int `json:"memory"`
				Disk   int `json:"disk"`
			} `json:"allocated_resources"`
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

type Node struct {
	Object     string `json:"object"`
	Attributes struct {
		Id                 int         `json:"id"`
		Uuid               string      `json:"uuid"`
		Public             bool        `json:"public"`
		Name               string      `json:"name"`
		Description        interface{} `json:"description"`
		LocationId         int         `json:"location_id"`
		Fqdn               string      `json:"fqdn"`
		Scheme             string      `json:"scheme"`
		BehindProxy        bool        `json:"behind_proxy"`
		MaintenanceMode    bool        `json:"maintenance_mode"`
		Memory             int         `json:"memory"`
		MemoryOverallocate int         `json:"memory_overallocate"`
		Disk               int         `json:"disk"`
		DiskOverallocate   int         `json:"disk_overallocate"`
		UploadSize         int         `json:"upload_size"`
		DaemonListen       int         `json:"daemon_listen"`
		DaemonSftp         int         `json:"daemon_sftp"`
		DaemonBase         string      `json:"daemon_base"`
		CreatedAt          time.Time   `json:"created_at"`
		UpdatedAt          time.Time   `json:"updated_at"`
		AllocatedResources struct {
			Memory int `json:"memory"`
			Disk   int `json:"disk"`
		} `json:"allocated_resources"`
	} `json:"attributes"`
}

type CreateNode struct {
	Object     string `json:"object"`
	Attributes struct {
		Id                 int         `json:"id"`
		Uuid               string      `json:"uuid"`
		Public             bool        `json:"public"`
		Name               string      `json:"name"`
		Description        interface{} `json:"description"`
		LocationId         int         `json:"location_id"`
		Fqdn               string      `json:"fqdn"`
		Scheme             string      `json:"scheme"`
		BehindProxy        bool        `json:"behind_proxy"`
		MaintenanceMode    bool        `json:"maintenance_mode"`
		Memory             int         `json:"memory"`
		MemoryOverallocate int         `json:"memory_overallocate"`
		Disk               int         `json:"disk"`
		DiskOverallocate   int         `json:"disk_overallocate"`
		UploadSize         int         `json:"upload_size"`
		DaemonListen       int         `json:"daemon_listen"`
		DaemonSftp         int         `json:"daemon_sftp"`
		DaemonBase         string      `json:"daemon_base"`
		CreatedAt          time.Time   `json:"created_at"`
		UpdatedAt          time.Time   `json:"updated_at"`
		AllocatedResources struct {
			Memory int `json:"memory"`
			Disk   int `json:"disk"`
		} `json:"allocated_resources"`
	} `json:"attributes"`
	Meta struct {
		Resource string `json:"resource"`
	} `json:"meta"`
}

type CreateNodeRequest struct {
	Name               string `json:"name"`
	LocationId         int    `json:"location_id"`
	Fqdn               string `json:"fqdn"`
	Scheme             string `json:"scheme"`
	Memory             int    `json:"memory"`
	MemoryOverallocate int    `json:"memory_overallocate"`
	Disk               int    `json:"disk"`
	DiskOverallocate   int    `json:"disk_overallocate"`
	UploadSize         int    `json:"upload_size"`
	DaemonSftp         int    `json:"daemon_sftp"`
	DaemonListen       int    `json:"daemon_listen"`
}

const nodesPath = "/application/nodes"

func (app *PterodactylConfig) GetNodes() (Nodes, error) {
	var nodes Nodes
	err := app.apiCall(nodesPath, "GET", nil, &nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

func (app *PterodactylConfig) GetNode(id int) (Node, error) {
	var node Node
	nodePath := fmt.Sprintf("%s/%d", nodesPath, id)
	err := app.apiCall(nodePath, "GET", nil, &node)
	if err != nil {
		return node, err
	}
	return node, nil
}

func (app *PterodactylConfig) CreateNode(node CreateNodeRequest) (CreateNode, error) {
	var result CreateNode
	jsonBody, err := json.Marshal(node)
	if err != nil {
		return result, err
	}
	err = app.apiCall(nodesPath, "POST", jsonBody, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (app *PterodactylConfig) UpdateNode(id int, node CreateNodeRequest) (CreateNode, error) {
	var result CreateNode
	jsonBody, err := json.Marshal(node)
	if err != nil {
		return result, err
	}
	err = app.apiCall(fmt.Sprintf("%s/%d", nodesPath, id), "PATCH", jsonBody, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (app *PterodactylConfig) DeleteNode(id int) error {
	err := app.apiCall(fmt.Sprintf("%s/%d", nodesPath, id), "DELETE", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
