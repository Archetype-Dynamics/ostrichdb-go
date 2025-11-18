package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ostrichdb-go/src/lib"
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

 //Builds a pointer to Collection of the given name (n)
 // Passed a pointer to a Project (p)
func NewCollectionBuilder (p *lib.Project, n string) *lib.Collection{
	return &lib.Collection{
		Project : p,
		Name: n,
	}
}

// Sends a POST request over the OstrichDB server
// to create a new Collection (c) in a Project
func CreateCollection(c *lib.Collection) error {
	client:= c.Project.Client
	pName:= c.Project.Name
	colName:= c.Name

	path:= lib.PathBuilder(pName, colName)

	response, err:=  lib.Post(client, path, "application/json", nil)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("Failed to create collection %s: in project: %s", colName, pName)
	}

	return nil
}

// Sends a GET request over the OstrichDB server
// to fetch a list of all Collections within a Project (p)
func ListCollections(p *lib.Project) error {
	client:= p.Client
	pName:= p.Name

	path:= lib.PathBuilder(pName)

	response, err:=  lib.Get(client, path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to get collections in project: %s ", pName)
	}

	return nil
}

// Sends a DELETE request over the OstrichDB server
// to remove Collection (c) from a Project
func DeleteCollection(c *lib.Collection) error {
	client:= c.Project.Client
	pName:= c.Project.Name
	colName:= c.Name

	path:= lib.PathBuilder(pName, colName)

	response, err:=  lib.Delete(client, path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to delete collection %s: in project: %s", colName, pName)
	}

	return nil
}

// Sends a PUT request over the OstrichDB server
// to rename a Collection (c) to (new)
func RenameCollection(c *lib.Collection, new string) error {
	client:= c.Project.Client
	pName:= c.Project.Name
	colName:= c.Name

	path:= fmt.Sprintf("%s/projects/%s/collections/%s?rename=%s", lib.OSTRICHDB_ADDRESS, pName, colName, new)

	response, err:=  lib.Put(client, path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to rename collection: %s in project: %s to %s", colName, pName, new)
	}

	return nil
}

// Sends a GET request over the OstrichDB server
// to fetch a Collection's (c) metadata
func GetCollectionInfo(c *lib.Collection) lib.CollectionInfo{
	var collectionInfo lib.CollectionInfo

	client:= c.Project.Client
	pName:= c.Project.Name
	colName:= c.Name

	path:= lib.PathBuilder(pName, colName)

	response, _:= lib.Get(client, path)
	defer response.Body.Close()

	resBody, _ := io.ReadAll(response.Body)
	unmarshalError:= json.Unmarshal(resBody, &collectionInfo)
	if unmarshalError != nil {
		return collectionInfo
	}

	return collectionInfo
}