package Pterodactyl

import (
	"encoding/json"
	"fmt"
	"time"
)

type Databases struct {
	Object string `json:"object"`
	Data   []struct {
		Object     string `json:"object"`
		Attributes struct {
			Id             int       `json:"id"`
			Server         int       `json:"server"`
			Host           int       `json:"host"`
			Database       string    `json:"database"`
			Username       string    `json:"username"`
			Remote         string    `json:"remote"`
			MaxConnections int       `json:"max_connections"`
			CreatedAt      time.Time `json:"created_at"`
			UpdatedAt      time.Time `json:"updated_at"`
			Relationships  struct {
				Password struct {
					Object     string `json:"object"`
					Attributes struct {
						Password string `json:"password"`
					} `json:"attributes"`
				} `json:"password"`
				Host struct {
					Object     string `json:"object"`
					Attributes struct {
						Id        int       `json:"id"`
						Name      string    `json:"name"`
						Host      string    `json:"host"`
						Port      int       `json:"port"`
						Username  string    `json:"username"`
						Node      int       `json:"node"`
						CreatedAt time.Time `json:"created_at"`
						UpdatedAt time.Time `json:"updated_at"`
					} `json:"attributes"`
				} `json:"host"`
			} `json:"relationships"`
		} `json:"attributes"`
	} `json:"data"`
}

type Database struct {
	Object     string `json:"object"`
	Attributes struct {
		Id             int       `json:"id"`
		Server         int       `json:"server"`
		Host           int       `json:"host"`
		Database       string    `json:"database"`
		Username       string    `json:"username"`
		Remote         string    `json:"remote"`
		MaxConnections int       `json:"max_connections"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	} `json:"attributes"`
	Meta struct {
		Resource string `json:"resource,omitempty"`
	} `json:"meta,omitempty"`
}

const databasePath = "/application/servers/%d/databases"

func (c *Client) GetDatabases(serverId int) (Databases, error) {
	var databases Databases
	path := fmt.Sprintf(databasePath, serverId)
	err := c.ApiCall( path, "GET", nil, &databases)
	if err != nil {
		return databases, err
	}
	return databases, nil
}

func (c *Client) GetDatabase(serverId, id int) (Database, error) {
	var database Database
	path := fmt.Sprintf("%s/%d", fmt.Sprintf(databasePath, serverId), id)
	err := c.ApiCall(path, "GET", nil, &database)
	if err != nil {
		return database, err
	}
	return database, nil
}

func (c *Client) CreateDatabase(serverId int, database Database) (Database, error) {
	var createdDatabase Database
	path := fmt.Sprintf(databasePath, serverId)
	body, err := json.Marshal(database)
	if err != nil {
		return createdDatabase, err
	}
	err = c.ApiCall(path, "POST", body, &createdDatabase)
	if err != nil {
		return createdDatabase, err
	}
	return createdDatabase, nil
}

func(c *Client) DeleteDatabase(serverId, id int) error {
	path := fmt.Sprintf("%s/%d", fmt.Sprintf(databasePath, serverId), id)
	err := c.ApiCall(path, "DELETE", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ResetDatabasePassword(serverId, id int) error {
	path := fmt.Sprintf("%s/%d/reset-password", fmt.Sprintf(databasePath, serverId), id)
	err := c.ApiCall(path, "POST", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
