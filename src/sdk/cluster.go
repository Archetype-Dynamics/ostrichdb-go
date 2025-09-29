package sdk

import "./lib"
import (
    "fmt"
    "net/http"
)


type Cluster struct{
	Client *Client
	Project *Project
	Collection *Collection
	Name string
}

func (c *Client) NewClusterBuilder(proj *Project, col *Collection, name string) *Cluster{
	return &Cluster{
		Client: c,
		Project: proj,
		Collection: col,
		Name: name,
	}
}

func (c *Client) CreateCluster (cluster *Cluster) error {
	projName:= cluster.Project.Name
	collectionName:= cluster.Collection.Name

	path:= fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s", lib.OSTRICHDB_ADDRESS, projName, collectionName, cluster.Name )

	response, err:= http.Post(path, "application/json", nil)
	if err != nil {
		return err
	}

	defer response.Body.Close()


	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to create cluster %s: in collection %s in project: %s", cluster.Name, collectionName, projName)
	}

	return nil
}

//Todo: continue on

