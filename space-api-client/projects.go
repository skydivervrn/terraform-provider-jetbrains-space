package space

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetProject - Returns project by id
func (c *Client) GetProject(id string) (Project, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/http/projects/id:%s", c.HostURL, id), nil)
	if err != nil {
		return Project{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return Project{}, err
	}

	project := Project{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return Project{}, err
	}

	return project, nil
}

//getProjectRepos
func (c *Client) getProjectRepos(projectId string) (ProjectRepos, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s/id:%s?$fields=repos", c.HostURL, baseApiEndpoint, projectId), nil)
	if err != nil {
		return ProjectRepos{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return ProjectRepos{}, err
	}

	project := ProjectRepos{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return ProjectRepos{}, err
	}

	return project, nil
}

// CreateProject - Creates project with given name
func (c *Client) CreateProject(name string) (Project, error) {
	data := new(struct {
		Key struct {
			Key string `json:"key"`
		} `json:"key"`
		Name string `json:"name"`
	})
	data.Name = name
	data.Key.Key = strings.ToUpper(strings.ReplaceAll(name, " ", "-"))
	bytesData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.HostURL, baseApiEndpoint), bytes.NewBuffer(bytesData))
	if err != nil {
		return Project{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return Project{}, err
	}

	project := Project{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return Project{}, err
	}

	return project, nil
}

// UpdateProject - Creates project with given name
func (c *Client) UpdateProject(id, name string) (Project, error) {
	data := new(struct {
		Name string `json:"name"`
	})
	data.Name = name
	bytesData, _ := json.Marshal(data)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s%s/id:%s", c.HostURL, baseApiEndpoint, id), bytes.NewBuffer(bytesData))
	if err != nil {
		return Project{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return Project{}, err
	}

	project := Project{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return Project{}, err
	}

	return project, nil
}

// DeleteProject - Creates project with given name
func (c *Client) DeleteProject(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s/id:%s", c.HostURL, baseApiEndpoint, id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
