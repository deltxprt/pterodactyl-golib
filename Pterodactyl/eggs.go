package Pterodactyl

import (
	"fmt"
	"time"
)

type Eggs struct {
	Object string `json:"object"`
	Data   []struct {
		Object     string `json:"object"`
		Attributes struct {
			Id           int         `json:"id"`
			Uuid         string      `json:"uuid"`
			Name         string      `json:"name"`
			Nest         int         `json:"nest"`
			Author       string      `json:"author"`
			Description  string      `json:"description"`
			DockerImage  string      `json:"docker_image"`
			DockerImages interface{} `json:"docker_images"`
			Config       interface{} `json:"config"`
			Startup      string      `json:"startup"`
			Script       struct {
				Privileged bool        `json:"privileged"`
				Install    string      `json:"install"`
				Entry      string      `json:"entry"`
				Container  string      `json:"container"`
				Extends    interface{} `json:"extends"`
			} `json:"script"`
			CreatedAt     time.Time `json:"created_at"`
			UpdatedAt     time.Time `json:"updated_at"`
			Relationships struct {
				Nest struct {
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
				} `json:"nest"`
				Servers struct {
					Object string `json:"object"`
					Data   []struct {
						Object     string `json:"object"`
						Attributes struct {
							Id          int         `json:"id"`
							ExternalId  interface{} `json:"external_id"`
							Uuid        string      `json:"uuid"`
							Identifier  string      `json:"identifier"`
							Name        string      `json:"name"`
							Description string      `json:"description"`
							Status      interface{} `json:"status"`
							Suspended   bool        `json:"suspended"`
							Limits      struct {
								Memory      int         `json:"memory"`
								Swap        int         `json:"swap"`
								Disk        int         `json:"disk"`
								Io          int         `json:"io"`
								Cpu         int         `json:"cpu"`
								Threads     interface{} `json:"threads"`
								OomDisabled bool        `json:"oom_disabled"`
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
								StartupCommand string      `json:"startup_command"`
								Image          string      `json:"image"`
								Installed      int         `json:"installed"`
								Environment    interface{} `json:"environment"`
							} `json:"container"`
							UpdatedAt time.Time `json:"updated_at"`
							CreatedAt time.Time `json:"created_at"`
						} `json:"attributes"`
					} `json:"data"`
				} `json:"servers"`
			} `json:"relationships"`
		} `json:"attributes"`
	} `json:"data"`
}

type Egg struct {
	Object     string `json:"object"`
	Attributes struct {
		Id           int         `json:"id"`
		Uuid         string      `json:"uuid"`
		Name         string      `json:"name"`
		Nest         int         `json:"nest"`
		Author       string      `json:"author"`
		Description  string      `json:"description"`
		DockerImage  string      `json:"docker_image"`
		DockerImages interface{} `json:"docker_images"`
		Config       interface{} `json:"config"`
		Startup      string      `json:"startup"`
		Script       struct {
			Privileged bool        `json:"privileged"`
			Install    string      `json:"install"`
			Entry      string      `json:"entry"`
			Container  string      `json:"container"`
			Extends    interface{} `json:"extends"`
		} `json:"script"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`
}

const eggPath = "/application/nests/%d/eggs"

func (c *Client) GetEggs(nest int) (Eggs, error) {
	var eggs Eggs
	path := fmt.Sprintf(eggPath, nest)
	err := c.ApiCall(path, "GET", nil, &eggs)
	if err != nil {
		return eggs, err
	}
	return eggs, nil
}

func (c *Client) GetEgg(nest, id int) (Egg, error) {
	var egg Egg
	path := fmt.Sprintf("%s/%d", fmt.Sprintf(eggPath, nest), id)
	err := c.ApiCall(path, "GET", nil, &egg)
	if err != nil {
		return egg, err
	}
	return egg, nil
}
