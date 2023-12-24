package Pterodactyl

import (
	"encoding/json"
	"fmt"
	"time"
)

type Servers struct {
	Object string `json:"object"`
	Data   []struct {
		Object     string `json:"object"`
		Attributes struct {
			Id          int    `json:"id"`
			ExternalId  int    `json:"external_id"`
			Uuid        string `json:"uuid"`
			Identifier  string `json:"identifier"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Status      string `json:"status"`
			Suspended   bool   `json:"suspended"`
			Limits      struct {
				Memory      int  `json:"memory"`
				Swap        int  `json:"swap"`
				Disk        int  `json:"disk"`
				Io          int  `json:"io"`
				Cpu         int  `json:"cpu"`
				Threads     int  `json:"threads"`
				OomDisabled bool `json:"oom_disabled"`
			} `json:"limits"`
			FeatureLimits struct {
				Databases   int `json:"databases"`
				Allocations int `json:"allocations"`
				Backups     int `json:"backups"`
			} `json:"feature_limits"`
			User       int `json:"user"`
			Node       int `json:"node"`
			Allocation int `json:"allocation"`
			Nest       int `json:"nest"`
			Egg        int `json:"egg"`
			Container  struct {
				StartupCommand string            `json:"startup_command"`
				Image          string            `json:"image"`
				Installed      int               `json:"installed"`
				Environment    map[string]string `json:"environment"`
			} `json:"container"`
			UpdatedAt time.Time `json:"updated_at"`
			CreatedAt time.Time `json:"created_at"`
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

type Server struct {
	Object     string `json:"object"`
	Attributes struct {
		Id          int    `json:"id"`
		ExternalId  int    `json:"external_id"`
		Uuid        string `json:"uuid"`
		Identifier  string `json:"identifier"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Status      string `json:"status,omitempty"`
		Suspended   bool   `json:"suspended"`
		Limits      struct {
			Memory      int  `json:"memory"`
			Swap        int  `json:"swap"`
			Disk        int  `json:"disk"`
			Io          int  `json:"io"`
			Cpu         int  `json:"cpu"`
			Threads     int  `json:"threads"`
			OomDisabled bool `json:"oom_disabled"`
		} `json:"limits"`
		FeatureLimits struct {
			Databases   int `json:"databases"`
			Allocations int `json:"allocations"`
			Backups     int `json:"backups"`
		} `json:"feature_limits"`
		User       int `json:"user"`
		Node       int `json:"node"`
		Allocation int `json:"allocation"`
		Nest       int `json:"nest"`
		Egg        int `json:"egg"`
		Container  struct {
			StartupCommand string            `json:"startup_command"`
			Image          string            `json:"image"`
			Installed      int               `json:"installed"`
			Environment    map[string]string `json:"environment"`
		} `json:"container"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"attributes"`
}

type UpdateServerRequest struct {
	Name        string `json:"name"`
	User        int    `json:"user"`
	ExternalId  string `json:"external_id"`
	Description string `json:"description"`
}

type UpdateServerBuildRequest struct {
	Allocation    int `json:"allocation"`
	Memory        int `json:"memory"`
	Swap          int `json:"swap"`
	Disk          int `json:"disk"`
	Io            int `json:"io"`
	Cpu           int `json:"cpu"`
	Threads       int `json:"threads,omitempty"`
	FeatureLimits struct {
		Databases   int `json:"databases"`
		Allocations int `json:"allocations"`
		Backups     int `json:"backups"`
	} `json:"feature_limits"`
}

type UpdateServerStartupRequest struct {
	Startup     string            `json:"startup"`
	Environment map[string]string `json:"environment"`
	Egg         int               `json:"egg"`
	Image       string            `json:"image"`
	SkipScripts bool              `json:"skip_scripts"`
}

type CreateServerRequest struct {
	Name        string            `json:"name"`
	User        int               `json:"user"`
	Egg         int               `json:"egg"`
	DockerImage string            `json:"docker_image"`
	Startup     string            `json:"startup"`
	Environment map[string]string `json:"environment"`
	Limits      struct {
		Memory int `json:"memory"`
		Swap   int `json:"swap"`
		Disk   int `json:"disk"`
		Io     int `json:"io"`
		Cpu    int `json:"cpu"`
	} `json:"limits"`
	FeatureLimits struct {
		Databases int `json:"databases"`
		Backups   int `json:"backups"`
	} `json:"feature_limits"`
	Allocation struct {
		Default int `json:"default"`
	} `json:"allocation"`
}

const serverPath = "/application/servers"

func GetServers(pterodactylCfg PterodactylConfig) (Servers, error) {
	var servers Servers
	err := ApiCall(pterodactylCfg, serverPath, "GET", nil, &servers)
	if err != nil {
		return servers, err
	}
	return servers, nil
}

func GetServer(pterodactylCfg PterodactylConfig, id int) (Server, error) {
	var server Server
	path := fmt.Sprintf("%s/%d", serverPath, id)
	err := ApiCall(pterodactylCfg, path, "GET", nil, &server)
	if err != nil {
		return server, err
	}
	return server, nil
}

func UpdateServer(pterodactylCfg PterodactylConfig, id int, request UpdateServerRequest) (Server, error) {
	var server Server
	path := fmt.Sprintf("%s/%d/details", serverPath, id)
	body, err := json.Marshal(request)
	if err != nil {
		return server, err
	}
	err = ApiCall(pterodactylCfg, path, "PATCH", body, &server)
	if err != nil {
		return server, err
	}
	return server, nil
}

func UpdateServerBuild(pterodactylCfg PterodactylConfig, id int, request UpdateServerBuildRequest) (Server, error) {
	var server Server
	path := fmt.Sprintf("%s/%d/build", serverPath, id)
	body, err := json.Marshal(request)
	if err != nil {
		return server, err
	}
	err = ApiCall(pterodactylCfg, path, "PATCH", body, &server)
	if err != nil {
		return server, err
	}
	return server, nil
}

func UpdateServerStartup(pterodactylCfg PterodactylConfig, id int, request UpdateServerStartupRequest) (Server, error) {
	var server Server
	path := fmt.Sprintf("%s/%d/startup", serverPath, id)
	body, err := json.Marshal(request)
	if err != nil {
		return server, err
	}
	err = ApiCall(pterodactylCfg, path, "PATCH", body, &server)
	if err != nil {
		return server, err
	}
	return server, nil
}

func CreateServer(pterodactylCfg PterodactylConfig, request CreateServerRequest) (Server, error) {
	var server Server
	body, err := json.Marshal(request)
	if err != nil {
		return server, err
	}
	err = ApiCall(pterodactylCfg, serverPath, "POST", body, &server)
	if err != nil {
		return server, err
	}
	return server, nil
}

func SuspendServer(pterodactylCfg PterodactylConfig, id int) error {
	path := fmt.Sprintf("%s/%d/suspend", serverPath, id)
	err := ApiCall(pterodactylCfg, path, "POST", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func UnsuspendServer(pterodactylCfg PterodactylConfig, id int) error {
	path := fmt.Sprintf("%s/%d/unsuspend", serverPath, id)
	err := ApiCall(pterodactylCfg, path, "POST", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func ReinstallServer(pterodactylCfg PterodactylConfig, id int) error {
	path := fmt.Sprintf("%s/%d/reinstall", serverPath, id)
	err := ApiCall(pterodactylCfg, path, "POST", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func DeleteServer(pterodactylCfg PterodactylConfig, id int) error {
	path := fmt.Sprintf("%s/%d", serverPath, id)
	err := ApiCall(pterodactylCfg, path, "DELETE", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func ForceDeleteServer(pterodactylCfg PterodactylConfig, id int) error {
	path := fmt.Sprintf("%s/%d/force", serverPath, id)
	err := ApiCall(pterodactylCfg, path, "DELETE", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
