package ostrichdb

import "./lib"
import (
    "fmt"
    "net/http"
)


type Collection struct {
    Client *Client
    ProjectName string
    Name string
}


func (c *Client) make_new_collection_builder (projName string, collectionName string) *Collection{
	return &Collection{
		Client: c,
		ProjectName: projName,
		Name: collectionName,
	}
}

//do any work on a collection just pass in the struct and method
func (c *client) some_function(method lib.HttpMethod, collection *Collection) error {

	path:= fmt.Sprintf("%s/projects/%s/collections/%s", lib.OSTRICHDB_ADDRESS, collection.ProjectName, collection.Name)

	var err error
	var response *http.Response
	var operationStr string

	switch(lib.HttpMethod){
		case lib.HTTP_GET:
			methodStr = "fetch"
			response, err = http.Get(path, "application/json", nil)
			break
		case lib.HTTP_POST:
		methodStr = "create"
			response, err = http.Post(path, "application/json", nil)
			break
		case lib.HTTP_DELETE:
		methodStr = "delete"
			response, err = lib.Delete(path)
			break
		case lib.HTTP_PUT:
		methodStr = "rename"
			response, err = lib.Put(path)
			break
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to %s collection %s: in project: %s", operationStr,collection.Name, collection.ProjectName)
	}

	return nil
}


func (c *Client) create_collection(collection *Collection) error {
	path:= fmt.Sprintf("%s/projects/%s/collections/%s", lib.OSTRICHDB_ADDRESS, collection.ProjectName, collection.Name)

	response, err:=  http.Post(path, "application/json", nil)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to create collection %s: in project: %s", collection.Name, collection.ProjectName)
	}

	return nil
}
