package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ostrichdb-go/src/lib"
	"strconv"

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

	path:= lib.PathBuilder(lib.QUERY_PARAM_NONE,pName, colName, cluName)

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

	path:= lib.PathBuilder(lib.QUERY_PARAM_NONE, pName, colName, cluName)

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
func FetchCluster(c *lib.Cluster) (*lib.Cluster, error){
	var cluster *lib.Cluster

	client:= c.Collection.Project.Client
	pName:= c.Collection.Project.Name
	colName:= c.Collection.Name
	cluName:= c.Name

	path:= lib.PathBuilder(lib.QUERY_PARAM_NONE,pName, colName, cluName)

	response, err:= lib.Get(client, path)
	if err != nil {
		return cluster, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return cluster, fmt.Errorf("Failed to fetch cluster %s: in collection %s in project: %s", cluName, colName, pName)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return cluster, fmt.Errorf("Failed to fetch Cluster of name: %s", cluName)
	}


	err = json.Unmarshal(data, &cluster)
	if err != nil {
		return cluster, err
	}

	return cluster, nil
}

// Sends a PUT request over the OstrichDB server
// to rename a Cluster (c) to (new)
func RenameCluster(c *lib.Cluster, new string) error{
	client:= c.Collection.Project.Client
	pName:= c.Collection.Project.Name
	colName:= c.Collection.Name
	cluName:= c.Name

	path:= lib.PathBuilder(lib.QUERY_PARAM_RENAME, pName, colName,cluName, new)

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


func ListClusters(c *lib.Collection) (lib.ClusterInfo, error){
	var clusterInfo lib.ClusterInfo

	client:= c.Project.Client
	pName:= c.Project.Name
	colName:= c.Name


	path:= lib.PathBuilder(lib.TIER, pName, colName)

	response, err:= lib.Get(client,path)

	if err != nil {
		return clusterInfo, err
	}

	if response.StatusCode != http.StatusOK{
		return clusterInfo, fmt.Errorf("Error getting cluster information from Collection named: %s", colName)
	}

	data, err:= io.ReadAll(response.Body)
	if err != nil || len(data) == 0 {
		return clusterInfo, err
	}

	err =  json.Unmarshal(data, &clusterInfo)
	if err != nil {
		return clusterInfo, err
	}

	return clusterInfo, nil
}

// This helper used to get a Collection's (c) info
// specifially the count of clusters within (c)
// converts to int and return
func GetClusterCount(c *lib.Collection) int{
	info := GetCollectionInfo(c)
 	countStr:= info.ClusterCount

  count, _:= strconv.Atoi(countStr)
  return count
}