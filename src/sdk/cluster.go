package sdk

import "../lib"
import (
    "fmt"
    "net/http"
)


func NewClusterBuilder(proj *lib.Project, col *lib.Collection, name string) *lib.Cluster{
	return &lib.Cluster{
		Client: proj.Client,
		Project: proj,
		Collection: col,
		Name: name,
	}
}

func CreateCluster (cluster *lib.Cluster) error {
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


func DeleteCluster (cluster *lib.Cluster) error {
	projName:= cluster.Project.Name
	collectionName:= cluster.Collection.Name
	path:= fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s", lib.OSTRICHDB_ADDRESS, projName, collectionName, cluster.Name )


	response, err:= lib.Delete(path)
	if err != nil {
		return err
	}

	defer response.Body.Close()


	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to delete cluster %s: in collection %s in project: %s", cluster.Name, collectionName, projName)
	}

	return nil
}

//Todo: Finish me
func ListClusters(collection *lib.Collection){}


func RenameCluster(cluster *lib.Cluster, new string) error{
	projName:= cluster.Project.Name
	collectionName:= cluster.Collection.Name
	path:= fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s?rename=%s", lib.OSTRICHDB_ADDRESS, projName, collectionName, cluster.Name, new)

	response, err:= lib.Put(path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to rename cluster: %s to %s in collection %s in project: %s", cluster.Name,new, collectionName, projName)
	}

	return nil
}