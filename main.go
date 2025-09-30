package main

import "fmt"
import "ostrichdb-go/src/sdk"


//Example for creating a new Collection(database)
func main(){

	//Mandatory: Config & Client setup
    config:= sdk.NewConfigBuilder()
    client:= sdk.NewClientBuilder(config)

    //Mandatory: Create a Project builder pass in desired project name.
    // Note: If a project already exists in you OstrichDB instance pass the name
    //  		of the already existing project
    project:= sdk.NewProjectBuilder(client, "myCoolProject")
    projectErr:= sdk.CreateProject(project)
    if projectErr != nil {
     	fmt.Println("Error creating project")
    }

    //Mandatory: Create a Collection builder pass in desired Collection name

    collection:= sdk.NewCollectionBuilder(project, "myFirstCollection")
    collectionErr:= sdk.CreateCollection(collection)

    if collectionErr != nil {
     	fmt.Println("Error creating collection")
    }
}