package sdk

import "ostrichdb-go/src/lib"
import (
    "fmt"
    "net/http"
)

func NewRecordBuilder(proj *lib.Project, col *lib.Collection, clu *lib.Cluster,name string,t string, value string) *lib.Record{
	return &lib.Record{
		Client: proj.Client,
		Project: proj,
		Collection: col,
		Cluster: clu,
		Name: name,
		Type: t,
		Value: value,
	}
}

func CreateRecord(record *lib.Record) error{
	projName:= record.Project.Name
	colName:= record.Collection.Name
	cluName:= record.Cluster.Name
	rName:= record.Name
	rType:= record.Type
	rVal:= record.Value

	path:= fmt.Sprintf("%s/projects/%s/collections/%s/clusters/%s/records/%s?type=%s&value=%s",
		lib.OSTRICHDB_ADDRESS,
	 	projName,
	 	colName,
	 	cluName,
		rName,
		rType,
		rVal)

	fmt.Println("Making a record at path: ", path)
	response, err:= lib.Post(record.Client, path, "application/json", nil)
	if err != nil {
		return err
	}

	defer response.Body.Close()


	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to create record %s: in cluster %s in collection %s in project: %s", rName, cluName, colName, projName)
	}

	return nil

}