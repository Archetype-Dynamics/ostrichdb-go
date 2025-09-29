package main

import (
	"fmt"
	"../lib"
	"../sdk"
)


//Example for creating a new Collection(database)
func main(){

	//Mandatory: Config & Client setup
    config:= sdk.NewConfigBuilder()
    client:= sdk.NewClientBuilder(config)

    //Mandatory: Create a Project builder pass in desired project name.
    // Note: If a project already exists in you OstrichDB instance pass the name
    //  		of the already existing project
    project:= (client).NewProjectBuilder("myCoolProject")

    //Mandatory: Create a Collection builder pass in desired Collection name

    collection:= (client).NewCollectionBuilder(project, "myFirstCollection")
    collectionErr:= (client).CreateCollection(collection)

    if collectionErr != nil {
     	//Handle error
    }
}