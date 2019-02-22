package api

import (
	"strconv"

	"github.com/pkg/errors"
)

type Application struct {
	ApplicationID   int    `json:"id"`
	Href            string `json:"href"`
	ApplicationName string `json:"application_name"`
}

type ApplicationCreateRequest struct {
	request
	Hostname  string `json:"hostname"`
	Origin    string `json:"origin"`
	StackName string `json:"stackName"`
}

type ApplicationCreateResponse struct {
	response
	ApplicationID   int    `json:"id"`
	Href            string `json:"href"`
	ApplicationName string `json:"application_name"`
}

func (c *client) ApplicationCreate(accountID int, hostname string, origin string, stackname string) (*ApplicationCreateResponse, error) {
	if accountID == 0 {
		return nil, errors.New("accountID parameter is required")
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

	req := &ApplicationCreateRequest{
		request:   c.newRequest("/account/" + strconv.Itoa(accountID) + "/application/create"),
		Hostname:  hostname,
		Origin:    origin,
		StackName: stackname,
	}

	var resp ApplicationCreateResponse
	err := c.httpPostJson("/account/"+strconv.Itoa(accountID)+"/application/create", req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "ApplicationCreate request failed.")
	}

	if resp.Message != "" {
		return nil, errors.New(resp.Message)
	}

	return &resp, nil
}
