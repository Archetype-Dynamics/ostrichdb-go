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

//Record Types enum see const.go
type RecordType int

//Record Types mapped to string val for storage. See const.go
var RecordTypeStrings = map[RecordType]string{
	NULL: "NULL",
	CHAR: "CHAR",
	STRING: "STRING",
	INTEGER: "INTEGER",
	FLOAT: "FLOAT",
	BOOLEAN: "BOOLEAN",
	DATE: "DATE",
	TIME: "TIME",
	DATETIME: "DATETIME",
	UUID: "UUID",
	CHAR_ARRAY: "[]CHAR",
	STRING_ARRAY: "[]STRING",
	INTEGER_ARRAY: "[]INTEGER",
	FLOAT_ARRAY: "[]FLOAT",
	BOOLEAN_ARRAY: "[]BOOLEAN",
	DATE_ARRAY: "[]DATE",
	TIME_ARRAY: "[]TIME",
	DATETIME_ARRAY: "[]DATETIME",
	UUID_ARRAY: "[]UUID",
}


//Represents the 3 common query params that are used
//when hitting OstrichDB endpoints.
//None: If there arent any...
//Type: When updating a Records type
//Value: When updating a Records Value
//Type_Value: When assigning a type AND value during Record creation
//Rename: For Collections, Clusters, and Records
type QueryType int
const (
	NONE QueryType = iota
	TYPE
	VALUE
	TYPE_VALUE
	RENAME

)