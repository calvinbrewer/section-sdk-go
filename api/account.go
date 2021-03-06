package api

import (
	"github.com/pkg/errors"
)

type accountApplication struct {
	ApplicationID   int    `json:"id"`
	Href            string `json:"href"`
	ApplicationName string `json:"application_name"`
}

type accountCreateRequest struct {
	request
	Name      string `json:"name"`
	Hostname  string `json:"hostname"`
	Origin    string `json:"origin"`
	StackName string `json:"stackName"`
}

type accountCreateResponse struct {
	response
	AccountID    int                  `json:"id"`
	Href         string               `json:"href"`
	AccountName  string               `json:"account_name"`
	Applications []accountApplication `json:"applications"`
}

func (c *client) AccountCreate(name string, hostname string, origin string, stackname string) (*accountCreateResponse, error) {
	if name == "" {
		return nil, errors.New("name parameter is required")
	}

	if hostname == "" {
		return nil, errors.New("hostname parameter is required")
	}

	if origin == "" {
		return nil, errors.New("origin parameter is required")
	}

	if stackname == "" {
		return nil, errors.New("stackname parameter is required")
	}

	req := &accountCreateRequest{
		request:   c.newRequest(),
		Name:      name,
		Hostname:  hostname,
		Origin:    origin,
		StackName: stackname,
	}

	var resp accountCreateResponse
	err := c.httpPostJson("/account/create", req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "AccountCreate request failed.")
	}

	if resp.Message != "" {
		return nil, errors.New(resp.Message)
	}

	return &resp, nil
}
