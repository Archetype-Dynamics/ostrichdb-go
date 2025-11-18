package lib

import (
	"fmt"
	"io"
	"net/http"
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


// Depending on the PathType(pt), Uses the passed in values (v)
// to construct a valid OstrichDB endpoint.
func PathBuilder(pt PathType,v ...string) string {
	switch(pt){
		case QUERY_PARAM_TYPE:
			if len(v) != 5{
				fmt.Printf("Incorrect number of values provided. Got %d expected 5\n", len(v))
				return ""
			}
			//Enusre the provided type is valid
			isValidDataType := false
			for _, dataType := range RecordTypeStrings {
				if v[4] == dataType {
					isValidDataType = true
					break
				}
			}
			if !isValidDataType {
				fmt.Printf("Invalid record type '%s' provided\n", v[5])
				return ""
			}
			return fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s/records/%s?type=%s", OSTRICHDB_ADDRESS, v[0], v[1], v[2], v[3], v[4])
		case QUERY_PARAM_VALUE:
			return fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s/records/%s?value=%s", OSTRICHDB_ADDRESS, v[0], v[1], v[2], v[3], v[4])
		case QUERY_PARAM_RENAME:
			switch(len(v)){
				case 3: //renaming a collection
					return fmt.Sprintf("%s/projects/%s/collections/%s?rename=%s", OSTRICHDB_ADDRESS, v[0], v[1], v[2])
				case 4: //renaming a cluster
					return fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s?rename=%s", OSTRICHDB_ADDRESS, v[0], v[1], v[2], v[3])
				case 5: //renaming a record
					return fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s/records/%s?rename=%s", OSTRICHDB_ADDRESS, v[0], v[1], v[2], v[3], v[4])
				default:
					fmt.Println("Invalid number of values provided while trying to rename. Defaulting")
					return ""
			}
		case QUERY_PARAM_TYPE_AND_VALUE:
			if len(v) != 6{
				fmt.Printf("Incorrect number of values provided. Got %d expected 6\n", len(v))
				return ""
			}
			isValidDataType:= false
			for _, dataType := range RecordTypeStrings{
				if v[4] == dataType {
					isValidDataType = true
					break
				}
			}
			if !isValidDataType {
				fmt.Printf("Invalid record type '%s' provided\n", v[5])
				return ""
			}
			return fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s/records/%s?type=%s&value=%s", OSTRICHDB_ADDRESS, v[0], v[1], v[2], v[3], v[4], v[5])
		case QUERY_PARAM_NONE:
			switch(len(v)){
				case 1:
					return fmt.Sprintf("%s/projects/%s", OSTRICHDB_ADDRESS, v[0])
				case 2:
					return fmt.Sprintf("%s/projects/%s/collections/%s", OSTRICHDB_ADDRESS, v[0], v[1])
				case 3:
					return fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s", OSTRICHDB_ADDRESS, v[0], v[1], v[2])
				case 4:
					return fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s/records/%s", OSTRICHDB_ADDRESS, v[0], v[1], v[2], v[3])
			}
	}

	return "Invalid number of names provided to PathBuilder()"
}