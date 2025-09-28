package ostrichdb

import "./lib"
import (
    "fmt"
    "net/http"
    "time"
)


type CollectionBuilder struct {
    Client *Client
    ProjectName string
    Name string
}


func (c *Client) make_new_collection_builder (projName string, collectionName string) *CollectionBuilder{
	return &CollectionBuilder{
		Client: c,
		ProjectName: projName,
		Name: collectionName,
	}
}


func (c *Client) create_collection(projName string, collectionName string) *error {

}