package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ostrichdb-go/src/lib"
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

	response, err:=  lib.Post(collection.Client, path, "application/json", nil)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("Failed to create collection %s: in project: %s", collection.Name, collection.ProjectName)
	}

	return nil
}

func ListCollections(project *lib.Project) error {
	path:= fmt.Sprintf("%s/projects/%s/collections", lib.OSTRICHDB_ADDRESS, project.Name)

	response, err:=  lib.Get(project.Client, path)
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

	response, err:=  lib.Delete(collection.Client, path)
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

	response, err:=  lib.Put(collection.Client, path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to rename collection: %s in project: &s to %s", collection.Name, collection.ProjectName, new)
	}

	return nil
}


type CollectionInfo struct {
	Name string `json:"name,omitempty"`
	ClusterCount  string `json:"cluster_count,omitempty"`
	RecordCount string `json:"record_count,omitempty"`
	Size string `json:"size,omitempty"`
}

func GetCollectionInfo(collection *lib.Collection) CollectionInfo{
	var collectionInfo CollectionInfo

	projName:= collection.ProjectName
	colName:= collection.Name

	path:= fmt.Sprintf("%s/projects/%s/collections/%s", lib.OSTRICHDB_ADDRESS, projName, colName)
	fmt.Println(path)
	response, _:= lib.Get(collection.Client, path)
	defer response.Body.Close()

	fmt.Println("Getting response.Header: ", response.Header)
	fmt.Println("Getting response.Body: ", response.Body)

	resBody, _ := io.ReadAll(response.Body)
	// fmt.Printf("data: %s\n",resBody)
	unmarshalError:= json.Unmarshal(resBody, &collectionInfo)
	if unmarshalError != nil {
		return collectionInfo
	}

	return collectionInfo
}
