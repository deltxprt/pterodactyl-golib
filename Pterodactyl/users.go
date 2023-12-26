package Pterodactyl

import (
	"encoding/json"
	"fmt"
	"time"
)

type Users struct {
	Object string `json:"object"`
	Data   []struct {
		Object     string `json:"object"`
		Attributes struct {
			Id         int         `json:"id"`
			ExternalId interface{} `json:"external_id"`
			Uuid       string      `json:"uuid"`
			Username   string      `json:"username"`
			Email      string      `json:"email"`
			FirstName  string      `json:"first_name"`
			LastName   string      `json:"last_name"`
			Language   string      `json:"language"`
			RootAdmin  bool        `json:"root_admin"`
			Fa         bool        `json:"2fa"`
			CreatedAt  time.Time   `json:"created_at"`
			UpdatedAt  time.Time   `json:"updated_at"`
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

type User struct {
	Object     string `json:"object"`
	Attributes struct {
		Id         int       `json:"id"`
		ExternalId string    `json:"external_id"`
		Uuid       string    `json:"uuid"`
		Username   string    `json:"username"`
		Email      string    `json:"email"`
		FirstName  string    `json:"first_name"`
		LastName   string    `json:"last_name"`
		Language   string    `json:"language"`
		RootAdmin  bool      `json:"root_admin"`
		Fa         bool      `json:"2fa"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	} `json:"attributes"`
	Meta struct {
		Resource string `json:"resource,omitempty"`
	} `json:"meta,omitempty"`
}

const usersPath = "/application/users"

func (c *Client) GetUsers() (Users, error) {
	var users Users
	err := c.ApiCall(usersPath, "GET", nil, &users)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (c *Client) GetUser(id int) (User, error) {
	var user User
	userPath := fmt.Sprintf("%s/%d", usersPath, id)
	err := c.ApiCall(userPath, "GET", nil, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (c *Client) CreateUser(username string, email string, firstName string, lastName string) (User, error) {
	var user User
	var body = map[string]string{
		"username":   username,
		"email":      email,
		"first_name": firstName,
		"last_name":  lastName,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return user, err
	}
	err = c.ApiCall(usersPath, "POST", jsonBody, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (c *Client) UpdateUser(id int, username string, email string, firstName string, lastName string, language string, password string) (User, error) {
	var user User
	var body = map[string]string{
		"email":      email,
		"username":   username,
		"first_name": firstName,
		"last_name":  lastName,
		"language":   language,
		"password":   password,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return user, err
	}
	userPath := fmt.Sprintf("%s/%d", usersPath, id)
	err = c.ApiCall(userPath, "PATCH", jsonBody, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (c *Client) DeleteUser(id int) error {
	userPath := fmt.Sprintf("%s/%d", usersPath, id)
	err := c.ApiCall(userPath, "DELETE", nil, nil)
	if err != nil {
		return err
	}
	return nil
}
