package sdk

import (
	"strconv"
	"fmt"
	"io"
	"net/http"
	"encoding/json"
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

//Builds a pointer to Cluster of the given name (n)
// Passed a pointer to a Collection (c)
func NewClusterBuilder(c *lib.Collection, n string) *lib.Cluster{
	return &lib.Cluster{
		Collection: c,
		Name: n,
	}
}

// Sends a POST request over the OstrichDB server
// to append a new Cluster (c) to a Collection
func CreateCluster (c *lib.Cluster) error {
	client:= c.Collection.Project.Client
	pName:= c.Collection.Project.Name
	colName:= c.Collection.Name
	cluName:= c.Name

	path:= lib.PathBuilder(pName, colName, cluName)

	response, err:= lib.Post(client, path, "application/json", nil)
	if err != nil {
		return err
	}

	defer response.Body.Close()


	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to create cluster %s: in collection %s in project: %s", cluName, colName, pName)
	}

	return nil
}

// Sends a DELETE request over the OstrichDB server
// to remove a Cluster (c) from a Collection
func DeleteCluster ( c *lib.Cluster) error {
	client:= c.Collection.Project.Client
	pName:= c.Collection.Project.Name
	colName:= c.Collection.Name
	cluName:= c.Name

	path:= lib.PathBuilder(pName, colName, cluName)

	response, err:= lib.Delete(client, path)
	if err != nil {
		return err
	}

	defer response.Body.Close()


	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to delete cluster %s: in collection %s in project: %s", cluName, colName, pName)
	}

	return nil
}

// Sends a GET request over the OstrichDB server
// to fetch all data(Records) from within a Cluster (c)
// TODO: KEep working on this
func FetchCluster(c *lib.Cluster) (string, error){
	client:= c.Collection.Project.Client
	pName:= c.Collection.Project.Name
	colName:= c.Collection.Name
	cluName:= c.Name

	path:= lib.PathBuilder(pName, colName, cluName)

	response, err:= lib.Get(client, path)
	if err != nil {
		return "GET Error", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "Status not ok error", fmt.Errorf("Failed to fetch cluster %s: in collection %s in project: %s", cluName, colName, pName)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "io Error", fmt.Errorf("ERROR getting content")
	}

	var cluster lib.Cluster
	unMarshallError:= json.Unmarshal(body, &cluster)
	if unMarshallError != nil {
		return "Unmarshal Error", unMarshallError
	}
	fmt.Println("CLUSTER: ", cluster)

	//TODO: instead of formating this to a string need to unmarshal this JSON into  a Record struct
	return fmt.Sprintf("%s", body), nil
}

// Sends a PUT request over the OstrichDB server
// to rename a Cluster (c) to (new)
func RenameCluster(c *lib.Cluster, new string) error{
	client:= c.Collection.Project.Client
	pName:= c.Collection.Project.Name
	colName:= c.Collection.Name
	cluName:= c.Name
	path:= fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s?rename=%s", lib.OSTRICHDB_ADDRESS, pName, colName, cluName, new)

	response, err:= lib.Put(client, path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to rename cluster: %s to %s in collection %s in project: %s", cluName,new, colName, pName)
	}

	return nil
}

//Todo: Finish me
func ListClusters(collection *lib.Collection){}

// This helper used to get a Collection's (c) info
// specifially the count of clusters within (c)
// converts to int and return
func GetClusterCount(c *lib.Collection) int{
	info := GetCollectionInfo(c)
 	countStr:= info.ClusterCount

  count, _:= strconv.Atoi(countStr)
  return count
}