package sdk

import "./lib"
import (
    "fmt"
    "net/http"
)


type Project struct {
	Client *Client
	Name string
}

func (c *Client) NewProjectBuilder(projName string) *Project{
	return &Project{
		Client: c,
		Name: projName,
	}
}

func (c *Client) create_project(proj *Project) error{
	projectPath := fmt.Sprintf("%s/projects/%s", lib.OSTRICHDB_ADDRESS, proj.Name)

	response, err := http.Post( projectPath, "application/json", nil)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to create project: %s, got status code: %d", proj.Name,response.StatusCode)
	}

	return nil
}

func (c *Client) delete_project(proj *Project) error {
	projectPath := fmt.Sprintf("%s/projects/%s", lib.OSTRICHDB_ADDRESS, proj.Name)

	response, err := lib.Delete(projectPath)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to delete project: %s, got status code: %d", proj.Name ,response.StatusCode)
	}

	return nil
}


func (c *Client) rename_project(proj *Project, new string) error{
	projectPath := fmt.Sprintf("%s/projects/%s?rename=%s", lib.OSTRICHDB_ADDRESS, proj.Name, new)

	response, err := lib.Put(projectPath)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to rename project: %s to %s got status code: %d", proj.Name, new, response.StatusCode)
	}

	return nil

}


func (c *Client) list_projects() ([]string, error){
	projectPath := fmt.Sprintf("%s/projects", lib.OSTRICHDB_ADDRESS)

	response, err := http.Get(projectPath)
	if err != nil {
		return []string{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []string{},fmt.Errorf("Failed to get project list, got status code: %d", response.StatusCode)
	}

	return  []string{}, nil
}