package sdk

import "./lib"
import (
    "fmt"
    "net/http"
)


type Record struct {
	Client *Client
	Project *Project
	Collection *Collection
	Cluster *Cluster
	Name string
	Type string
	Value string
}


func (c *Client) NewRecordBuilder(proj *Project, col *Collection, clu *Cluster, name string) *Record{
	return &Record{
		Client: c,
		Project: proj,
		Collection: col,
		Cluster: clu,
		Name: name,
	}
}