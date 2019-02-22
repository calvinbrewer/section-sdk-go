package api

import (
	"fmt"
	"github.com/pkg/errors"
)

type environmentCreateRequest struct {
	request
	Name                  string `json:"name"`
	SourceEnvironmentName string `json:"source_environment_name"`
	DomainName            string `json:"domain_name"`
}

type environmentDomain struct {
	Name     string `json:"name"`
	ZoneName string `json:"zoneName"`
	Cname    string `json:"cname"`
	Mode     string `json:"mode"`
}

type environmentCreateResponse struct {
	response
	EnvironmentID    int                 `json:"id"`
	Href             string              `json:"href"`
	EnvironmentName  string              `json:"environment_name"`
	Domains          []environmentDomain `json:"domains"`
	DNSBypassAddress string              `json:"dns_bypass_address"`
}

func (c *client) EnvironmentCreate(accountID int, applicationID int, name, sourceEnvironmentName, domainName string) (*environmentCreateResponse, error) {
	if accountID == 0 {
		return nil, errors.New("accountID parameter is required")
	}

	if applicationID == 0 {
		return nil, errors.New("applicationID parameter is required")
	}

	if name == "" {
		return nil, errors.New("name parameter is required")
	}

	if sourceEnvironmentName == "" {
		return nil, errors.New("sourceEnvironmentName parameter is required")
	}

	if domainName == "" {
		return nil, errors.New("domainName parameter is required")
	}

	req := &environmentCreateRequest{
		request: c.newRequest(),
		Name:    name,
		SourceEnvironmentName: sourceEnvironmentName,
		DomainName:            domainName,
	}

	var resp environmentCreateResponse
	err := c.httpPostJson(fmt.Sprintf("/account/%d/application/%d/environment/create", accountID, applicationID), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "AccountCreate request failed.")
	}

	if resp.Message != "" {
		return nil, errors.New(resp.Message)
	}

	return &resp, nil
}
