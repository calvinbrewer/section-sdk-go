package api

import (
	"github.com/pkg/errors"
)

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
}

type accountGetRequest struct {
	request
	AccountID int `json:"id"`
}

type accountGetResponse struct {
	response
	Account *Account `json:"account"`
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