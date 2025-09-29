package sdk

import "../lib"
import (
    "net/http"
    "time"
)


func NewConfigBuilder () *lib.Config{
	return &lib.Config{
		BaseURL: lib.OSTRICHDB_ADDRESS,
		Timeout: 30_000,
	}
}

func NewClientBuilder(config *lib.Config) *lib.Client {
	return &lib.Client{
		BaseURL: config.BaseURL,
		ApiKey:  config.ApiKey,
		HTTPClient: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Millisecond,
		},
	}
}

