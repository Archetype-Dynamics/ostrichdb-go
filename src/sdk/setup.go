package sdk

import "./lib"
import (
    "fmt"
    "net/http"
    "time"
)


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

func NewConfigBuilder () Config{
	config := new(Config)
	config.BaseURL = lib.OSTRICHDB_ADDRESS
	config.Timeout = 30_000

	return *config
}

func NewClientBuilder(config Config) *Client {
	return &Client{
		BaseURL: config.BaseURL,
		ApiKey:  config.ApiKey,
		HTTPClient: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Millisecond,
		},
	}
}

