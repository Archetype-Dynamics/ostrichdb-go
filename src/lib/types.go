package lib

import "net/http"
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

// Configuration settings for connecting to an OstrichDB backend.
type Config struct {
	BaseURL string `json:"baseUrl"`
	ApiKey string `json:"apiKey"`
	Timeout int `json:"timeout"`
}

// The main HTTP client for interacting with the OstrichDB API.
type Client struct {
	BaseURL string
	ApiKey string
	HTTPClient *http.Client
}

// The top-level unit of data in OstrichDB.
type Project struct {
	Client *Client
	Name string
}

// A container for clusters within a project.
type Collection struct {
    Project *Project
    Name string
}

// Grouping of related records within a collection.
type Cluster struct{
	Collection *Collection
	Name string
}

// An individual data entry within a cluster. Must have a name, type AND value
type Record struct {
	Cluster *Cluster
	Name string
	Type string
	Value string
}

// Holds metadata and statistics about a collection.
type CollectionInfo struct {
	Name string `json:"name,omitempty"`
	ClusterCount  string `json:"cluster_count,omitempty"`
	RecordCount string `json:"record_count,omitempty"`
	Size string `json:"size,omitempty"`
}