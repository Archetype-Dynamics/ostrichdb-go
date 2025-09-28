package ostrichdb

import "./lib"
import (
    "fmt"
    "net/http"
    "time"
)



type ProjectBuilder struct {
	Client *Client
	Name string
}

type Client struct {
	BaseURL string
	ApiKey string
	HTTPClient *http.Client
}

type Config struct {
	BaseURL string `json:"baseUrl"`
	ApiKey string `json:"apiKey"`
	Timeout int `json:"timeout"`
}

func make_new_config () Config{
	var config = new(Config)
	config.BaseURL = lib.OSTRICHDB_ADDRESS
	config.Timeout = 30_000

	return *config
}

func make_new_client(config Config) *Client {
	return &Client{
		BaseURL: config.BaseURL,
		ApiKey:  config.ApiKey,
		HTTPClient: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Millisecond,
		},
	}
}


func (c *Client) make_new_project_builder(projName string) *ProjectBuilder{
	return &ProjectBuilder {
		Client: c,
		Name: projName,
	}
}

func (c *Client) create_project(projName string) error{
	projectPath := fmt.Sprintf("%s/projects/%s", lib.OSTRICHDB_ADDRESS, projName)

	response, err := http.Post( projectPath, "application/json", nil)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to create project: %s, got status code: %d", projName,response.StatusCode)
	}

	return nil
}

func (c *Client) delete_project(projName string) error {
	projectPath := fmt.Sprintf("%s/projects/%s", lib.OSTRICHDB_ADDRESS, projName)

	response, err := lib.Delete(projectPath)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to delete project: %s, got status code: %d", projName,response.StatusCode)
	}

	return nil
}

//TODO: add a rename project function

func (c *Client) list_projects() ([]string, error){
	projectPath := fmt.Sprintf("%s/projects", lib.OSTRICHDB_ADDRESS)

	response, err := http.Get(projectPath)
	if err != nil {
		return []string{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []string{},fmt.Errorf("Failed to get project list, got status code: %d", response.StatusCode)
	}

	return  []string{}, nil
}