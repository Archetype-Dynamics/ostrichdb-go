package sdk

import "ostrichdb-go/src/lib"
import (
    "fmt"
    "net/http"
)


func NewCollectionBuilder (proj *lib.Project, collectionName string) *lib.Collection{
	return &lib.Collection{
		Client: proj.Client,
		ProjectName: proj.Name,
		Name: collectionName,
	}
}

func CreateCollection(collection *lib.Collection) error {
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

func ListCollections(project *lib.Project) error {
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


func DeleteCollection(collection *lib.Collection) error {
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


func RenameCollection(collection *lib.Collection, new string) error {
	path:= fmt.Sprintf("%s/projects/%s/collections/%s?rename=%s", lib.OSTRICHDB_ADDRESS, collection.ProjectName, collection.Name, new)

	response, err:=  lib.Put(path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to rename collection: %s in project: &s to %s", collection.Name, collection.ProjectName, new)
	}

	return nil
}