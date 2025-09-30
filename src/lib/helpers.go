package lib

import (
	"net/http"
	"io"
)

func Delete(client *Client, path string) (*http.Response, error) {
    request, err := http.NewRequest("DELETE", path, nil)
    if err != nil {
        return nil, err
    }
    request.Header.Set("Authorization", "Bearer " + client.ApiKey)
    return client.HTTPClient.Do(request)
}

func Put(client *Client, path string) (*http.Response, error) {
    request, err := http.NewRequest("PUT", path, nil)
    if err != nil {
        return nil, err
    }
    request.Header.Set("Authorization", "Bearer " + client.ApiKey)
    return client.HTTPClient.Do(request)
}

func Post(client *Client, path string, contentType string, body io.Reader) (*http.Response, error) {
    request, err := http.NewRequest("POST", path, body)
    if err != nil {
        return nil, err
    }
    request.Header.Set("Authorization", "Bearer " + client.ApiKey)
    request.Header.Set("Content-Type", contentType)
    return client.HTTPClient.Do(request)
}

func Get(client *Client, path string) (*http.Response, error) {
    request, err := http.NewRequest("GET", path, nil)
    if err != nil {
        return nil, err
    }
    request.Header.Set("Authorization", "Bearer " + client.ApiKey)
    return client.HTTPClient.Do(request)
}