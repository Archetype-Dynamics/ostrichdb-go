package sdk

import "ostrichdb-go/src/lib"
import "github.com/joho/godotenv"
import (
    "net/http"
    "time"
    "os"
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

//Gets the OSTRICHDB_JWT token from you env.
//Assumes its stored in a .env file by default
func GetAPIKey() string{
	godotenv.Load()
	ostrichJWT:= os.Getenv("OSTRICHDB_JWT")
	return ostrichJWT
}

//Builds a new default config
func NewConfigBuilder () *lib.Config{
	return &lib.Config{
		BaseURL: lib.OSTRICHDB_ADDRESS,
		Timeout: 30_000,
		ApiKey: GetAPIKey(),
	}
}

//Builds a new default client. Passed a pointer to a config
func NewClientBuilder(c *lib.Config) *lib.Client {
	return &lib.Client{
		BaseURL: c.BaseURL,
		ApiKey:  c.ApiKey,
		HTTPClient: &http.Client{
			Timeout: time.Duration(c.Timeout) * time.Millisecond,
		},
	}
}

