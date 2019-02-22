package api

import (
	"github.com/pkg/errors"
)

type Application struct {
	ApplicationID   int    `json:"id"`
	Href            string `json:"href"`
	ApplicationName string `json:"application_name"`
}

type Account struct {
	AccountID   int    `json:"id"`
	Href        string `json:"href"`
	AccountName string `json:"account_name"`
	Requires2FA string `json:"requires_2fa"`
	BillingUser string `json:"billing_user"`
	Owner       struct {
		ID          int    `json:"id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		Verified    bool   `json:"verified"`
		CompanyName string `json:"company_name"`
		PhoneNumber string `json:"phone_number"`
	} `json:"owner"`
	Applications []Application `json:"applications"`
}

type accountGetRequest struct {
	request
	AccountID int `json:"id"`
}

type accountGetResponse struct {
	response
	Account *Account `json:"account"`
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
	Account Account
}

func (c *client) AccountGet(AccountID int) (*Account, error) {
	if AccountID == 0 {
		return nil, errors.New("AccountID parameter is required.")
	}

	req := &accountGetRequest{
		request:   c.newRequest("Account.get"),
		AccountID: AccountID,
	}

	var resp accountGetResponse
	err := c.httpPostJson(req, &resp)
	if err != nil {
		return nil, errors.Wrapf(err, "AccountGet request failed for Account Id '%d'.", AccountID)
	}

	if resp.Code != "OK" || resp.Account == nil {
		return nil, newAPIError(resp.response, nil)
	}

	return resp.Account, nil
}

func (c *client) AccountCreate(name string, hostname string, origin string, stackname string) (*Account, error) {
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
		request:   c.newRequest("/account/create"),
		Name:      name,
		Hostname:  hostname,
		Origin:    origin,
		StackName: stackname,
	}

	var resp accountCreateResponse
	err := c.httpPostJson(req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "AccountCreate request failed.")
	}

	if resp.Message != "" {
		return nil, errors.New(resp.Message)
	}

	return &resp.Account, nil
}
