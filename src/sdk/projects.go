package sdk

import "ostrichdb-go/src/lib"
import (
	"net/http"
    "fmt"
)
/*
 *  Author: Marshall A Burns
 *  GitHub: @SchoolyB
 *
 *  Copyright (c) 2025-Present Archetype Dynamics, Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

// n = name of the Project
func NewProjectBuilder(c *lib.Client ,n string) *lib.Project{
	return &lib.Project{
		Client: c,
		Name: n,
	}
}

func CreateProject(p *lib.Project) error{
	pName:= p.Name

	path:= lib.PathBuilder(lib.NONE,pName)

	response, err := lib.Post(p.Client, path, "application/json", nil)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("Failed to create project: %s, got status code: %d", pName,response.StatusCode)
	}

	return nil
}

func DeleteProject(p *lib.Project) error {
	pName:= p.Name

	path:= lib.PathBuilder(lib.NONE, pName)

	response, err := lib.Delete(p.Client, path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to delete project: %s, got status code: %d", pName ,response.StatusCode)
	}

	return nil
}


func RenameProject(p *lib.Project, new string) error{
	pName:= p.Name

	path := fmt.Sprintf("%s/projects/%s?rename=%s", lib.OSTRICHDB_ADDRESS, pName, new)

	response, err := lib.Put(p.Client, path)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to rename project: %s to %s got status code: %d", pName, new, response.StatusCode)
	}

	return nil
}


func ListProjects(c *lib.Client) ([]string, error){
	path := fmt.Sprintf("%s/projects", lib.OSTRICHDB_ADDRESS)

	response, err := lib.Get(c, path)
	if err != nil {
		return []string{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []string{},fmt.Errorf("Failed to get project list, got status code: %d", response.StatusCode)
	}

	return  []string{}, nil
}