package sdk

import "ostrichdb-go/src/lib"
import (
    // "fmt"
    // "net/http"
)

func NewRecordBuilder(proj *lib.Project, col *lib.Collection, clu *lib.Cluster, name string) *lib.Record{
	return &lib.Record{
		Client: proj.Client,
		Project: proj,
		Collection: col,
		Cluster: clu,
		Name: name,
	}
}

// func CreateRecord(record *lib.Record) error{
// 	// project:= record.
// }