package sdk

import "ostrichdb-go/src/lib"
import (
    "fmt"
    "net/http"
)



func NewProjectBuilder(c *lib.Client ,projName string) *lib.Project{
	return &lib.Project{
		Client: c,
		Name: projName,
	}
}

func CreateProject(proj *lib.Project) error{
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

func DeleteProject(proj *lib.Project) error {
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


func RenameProject(proj *lib.Project, new string) error{
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


func ListProjects() ([]string, error){
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