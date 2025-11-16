package sdk

import "ostrichdb-go/src/lib"
import (
    "fmt"
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

 //Builds a pointer to Record of the given name (n) within a Cluster (clu)
 // with the type (t) and the value (v)
func NewRecordBuilder(clu *lib.Cluster,n string,t string, v string) *lib.Record{
	return &lib.Record{
		Cluster: clu,
		Name: n,
		Type: t,
		Value: v,
	}
}

// Sends a POST request over the OstrichDB server
// to create a new Record (r) in a Cluster
func CreateRecord(r *lib.Record) error{
	client:= r.Cluster.Collection.Project.Client
	pName:= r.Cluster.Collection.Project.Name
	colName:= r.Cluster.Collection.Name
	cluName:= r.Cluster.Name
	rName:= r.Name
	rType:= r.Type
	rVal:= r.Value

	path:= fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s/records/%s?type=%s&value=%s",
		lib.OSTRICHDB_ADDRESS,
	 	pName,
	 	colName,
	 	cluName,
		rName,
		rType,
		rVal)

	response, err:= lib.Post(client, path, "application/json", nil)
	if err != nil {
		return err
	}

	defer response.Body.Close()


	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to create record %s: in cluster %s in collection %s in project: %s", rName, cluName, colName, pName)
	}

	return nil
}

// Sends a DELETE request over the OstrichDB server
// to remove a Record (r) from a Cluster
func DeleteRecord(r *lib.Record)	error {
	client:= r.Cluster.Collection.Project.Client
	pName:= r.Cluster.Collection.Project.Name
	colName:= r.Cluster.Collection.Name
	cluName:= r.Cluster.Name
	rName:= r.Name

	path := fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s/records/%s", lib.OSTRICHDB_ADDRESS, pName, colName, cluName ,rName)

	response, err:= lib.Delete(client, path)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to delete record %s: in cluster %s within collection %s in project: %s", rName, cluName, colName, pName)
	}

	return nil
}


// Sends a PUT request over the OstrichDB server
// to rename a Record (r) to (new)
func RenameRecord(r *lib.Record, new string)error{
	client:= r.Cluster.Collection.Project.Client
	pName:= r.Cluster.Collection.Project.Name
	colName:= r.Cluster.Collection.Name
	cluName:= r.Cluster.Name
	rName:= r.Name

	path:= fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s/records/%s?rename=%s", lib.OSTRICHDB_ADDRESS, pName, colName, cluName, rName, new)

	response, err:= lib.Put(client, path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to rename record: %s to %s in cluster %s in collection %s in project: %s", rName, new, cluName, colName, pName)
	}

	return nil
}

// Sends a PUT request over the OstrichDB server
// to update a Record's (r) record type (t)
// For record types see types.go and const.go
func UpdateRecordType(r *lib.Record, t lib.RecordType)error{
	client:= r.Cluster.Collection.Project.Client
	pName:= r.Cluster.Collection.Project.Name
	colName:= r.Cluster.Collection.Name
	cluName:= r.Cluster.Name
	rName:= r.Name
	rType:= lib.RecordTypeStrings[t]

	path:= fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s/records/%s?type=%s", lib.OSTRICHDB_ADDRESS, pName, colName, cluName, rName, rType )

	response, err:= lib.Put(client, path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to update record: %s's type to %s in cluster %s in collection %s in project: %s", rName, rType, cluName, colName, pName)
	}

	return nil
}

// Sends a PUT request over the OstrichDB server
// to update a Record's (r) record value (v)
func UpdateRecordValue(r *lib.Record, v string) error {
	client:= r.Cluster.Collection.Project.Client
	pName:= r.Cluster.Collection.Project.Name
	colName:= r.Cluster.Collection.Name
	cluName:= r.Cluster.Name
	rName:= r.Name
	rValue:= v

	path:= fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s/records/%s?value=%s", lib.OSTRICHDB_ADDRESS, pName, colName, cluName, rName, rValue )

	response, err:= lib.Put(client, path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to update record: %s's value to %s in cluster %s in collection %s in project: %s", rName, rValue, cluName, colName, pName)
	}

	return nil
}