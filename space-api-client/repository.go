package space

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetRepository - Returns Repository by id
func (c *Client) GetRepository(repositoryName, projectId string) (Repository, error) {
	projectRepos, err := c.getProjectRepos(projectId)
	if err != nil {
		return Repository{}, err
	}
	for _, repo := range projectRepos.Repos {
		if repo.Name == repositoryName {
			return repo, nil
		}
	}
	return Repository{}, fmt.Errorf("repository %s not found", repositoryName)
}

// CreateRepository - Creates Repository with given name
func (c *Client) CreateRepository(repositoryName string, projectId string, data CreateRepositoryData) (Repository, error) {
	bytesData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s/id:%s/repositories/%s", c.HostURL, baseApiEndpoint, projectId, repositoryName), bytes.NewBuffer(bytesData))
	if err != nil {
		return Repository{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return Repository{}, err
	}

	repository := Repository{}
	err = json.Unmarshal(body, &repository)
	if err != nil {
		return Repository{}, err
	}

	return repository, nil
}

// UpdateRepository - Creates Repository with given name
//func (c *Client) UpdateRepository(id, name string) (Repository, error) {
//	data := new(struct {
//		Name string `json:"name"`
//	})
//	data.Name = name
//	bytesData, _ := json.Marshal(data)
//	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s%s/id:%s", c.HostURL, baseApiEndpoint, id), bytes.NewBuffer(bytesData))
//	if err != nil {
//		return Repository{}, err
//	}
//
//	body, err := c.doRequest(req)
//	if err != nil {
//		return Repository{}, err
//	}
//
//	Repository := Repository{}
//	err = json.Unmarshal(body, &Repository)
//	if err != nil {
//		return Repository{}, err
//	}
//
//	return Repository, nil
//}

// DeleteRepository - Creates Repository with given name
func (c *Client) DeleteRepository(projectId, repositoryName string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s/id:%s/repositories/%s", c.HostURL, baseApiEndpoint, projectId, repositoryName), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
