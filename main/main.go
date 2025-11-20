package main

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

import (
	"fmt"
	"ostrichdb-go/src/lib"
	"ostrichdb-go/src/sdk"
)

// Example demonstrating the full OstrichDB SDK workflow
// This creates a project, collection, cluster, and records
func main() {
	// Step 1: Config & Client setup
	// This loads your JWT token from the .env file
	config := sdk.NewConfigBuilder()
	client := sdk.NewClientBuilder(config)

	// Step 2: Create a Project
	// Projects are top-level containers in OstrichDB
	// Note: If the project already exists, simply use its existing name
	project := sdk.NewProjectBuilder(client, "exampleProject")
	projectErr := sdk.CreateProject(project)
	if projectErr != nil {
		fmt.Println("Error creating project:", projectErr)
		return
	}
	fmt.Println("✓ Project created successfully")

	// Step 3: Create a Collection
	// Collections (databases) belong to projects
	collection := sdk.NewCollectionBuilder(project, "users")
	collectionErr := sdk.CreateCollection(collection)
	if collectionErr != nil {
		fmt.Println("Error creating collection:", collectionErr)
		return
	}
	fmt.Println("✓ Collection created successfully")

	// Step 4: Create a Cluster
	// Clusters are groups of related records within a collection
	cluster := sdk.NewClusterBuilder(collection, "user_001")
	clusterErr := sdk.CreateCluster(cluster)
	if clusterErr != nil {
		fmt.Println("Error creating cluster:", clusterErr)
		return
	}
	fmt.Println("✓ Cluster created successfully")

	// Step 5: Create Records
	// Records store the actual data within clusters
	// Each record has a name, type, and value

	// Create a username record
	usernameRecord := sdk.NewRecordBuilder(cluster, "username", lib.STRING, "johndoe")
	if err := sdk.CreateRecord(usernameRecord); err != nil {
		fmt.Println("Error creating username record:", err)
		return
	}
	fmt.Println("✓ Username record created")

	// Create an email record
	emailRecord := sdk.NewRecordBuilder(cluster, "email", lib.STRING, "john@example.com")
	if err := sdk.CreateRecord(emailRecord); err != nil {
		fmt.Println("Error creating email record:", err)
		return
	}
	fmt.Println("✓ Email record created")

	// Create an age record (integer type)
	ageRecord := sdk.NewRecordBuilder(cluster, "age,", lib.INTEGER , "30")
	if err := sdk.CreateRecord(ageRecord); err != nil {
		fmt.Println("Error creating age record:", err)
		return
	}
	fmt.Println("✓ Age record created")

	// Step 6: Get cluster count
	clusterCount := sdk.GetClusterCount(collection)
	fmt.Printf("\n📊 Total clusters in collection: %d\n", clusterCount)

	// Step 7: Rename a cluster (optional)
	renameErr := sdk.RenameCluster(cluster, "user_001_renamed")
	if renameErr != nil {
		fmt.Println("Error renaming cluster:", renameErr)
		return
	}
	fmt.Println("✓ Cluster renamed successfully")

	// Step 8: Delete a record (optional)
	deleteErr := sdk.DeleteRecord(ageRecord)
	if deleteErr != nil {
		fmt.Println("Error deleting record:", deleteErr)
		return
	}
	fmt.Println("✓ Age record deleted successfully")

	fmt.Println("\n✅ All operations completed successfully!")
}
