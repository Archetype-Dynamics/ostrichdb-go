package lib
import (
	"net/http"
	"io"
)
/*
 *  Author: Marshall A Burns
 *  GitHub: @SchoolyB
 *
 *  Copyright (c) 2025-Present Archetype Dynamics, Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

//Helper function used to send a DELETE HTTP request to the OstrichDB server
// Args: 'c' is a pointer to a Client and 'p' is a path
func Delete(c *Client, p string) (*http.Response, error) {
    req, err := http.NewRequest("DELETE", p, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer " + c.ApiKey)
    return c.HTTPClient.Do(req)
}

//Helper function used to send a PUT HTTP request to the OstrichDB server
// Args: 'c' is a pointer to a Client and 'p' is a path
func Put(c *Client, p string) (*http.Response, error) {
    req, err := http.NewRequest("PUT", p, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer " + c.ApiKey)
    return c.HTTPClient.Do(req)
}

//Helper function used to send a POST HTTP request to the OstrichDB server
// Args: 'c' is a pointer to a Client and 'p' is a path
// 				'cType' or contentType is the type that is passed and set into the HTTP request header. e.g: "application/json"
// 				'body' is just passed nil when called
func Post(c *Client, p string, cType string, body io.Reader) (*http.Response, error) {
    req, err := http.NewRequest("POST", p, body)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer " + c.ApiKey)
    req.Header.Set("Content-Type", cType)
    return c.HTTPClient.Do(req)
}

//Helper function used to send a GET HTTP request to the OstrichDB server
// Gets API keys from a .env file and sets it in headers for auth
// Args: 'c' is a pointer to a Client and 'p' is a path
func Get(c *Client, p string) (*http.Response, error) {
    req, err := http.NewRequest("GET", p, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer " + c.ApiKey)
    return c.HTTPClient.Do(req)
}