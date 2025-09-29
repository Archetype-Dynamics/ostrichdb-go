package sdk

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


func (c *Client) NewCollectionBuilder (proj *Project, collectionName string) *Collection{
	return &Collection{
		Client: c,
		ProjectName: proj.Name,
		Name: collectionName,
	}
}

func (c *Client) CreateCollection(collection *Collection) error {
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

func (c *Client) ListCollections(project *Project) error {
	path:= fmt.Sprintf("%s/projects/%s/collections", lib.OSTRICHDB_ADDRESS, project.Name)

	response, err:=  http.Get(path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to get collections in project: %s ", project.Name)
	}

	return nil
}


func (c *Client) DeleteCollection(collection *Collection) error {
	path:= fmt.Sprintf("%s/projects/%s/collections/%s", lib.OSTRICHDB_ADDRESS, collection.ProjectName, collection.Name)

	response, err:=  lib.Delete(path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to delete collection %s: in project: %s", collection.Name, collection.ProjectName)
	}

	return nil
}


func (c *Client) RenameCollection(collection *Collection, new string) error {
	path:= fmt.Sprintf("%s/projects/%s/collections/%s?rename=%s", lib.OSTRICHDB_ADDRESS, collection.ProjectName, collection.Name, new)

	response, err:=  lib.Put(path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to rename collection: %s to %s", collection.Name, collection.ProjectName, new)
	}

	return nil
}