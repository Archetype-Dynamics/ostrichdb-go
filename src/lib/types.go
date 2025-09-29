package lib

import "net/http"

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

type Project struct {
	Client *Client
	Name string
}

type Collection struct {
    Client *Client
    ProjectName string
    Name string
}

type Cluster struct{
	Client *Client
	Project *Project
	Collection *Collection
	Name string
}

type Record struct {
	Client *Client
	Project *Project
	Collection *Collection
	Cluster *Cluster
	Name string
	Type string
	Value string
}