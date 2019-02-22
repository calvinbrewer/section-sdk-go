package api

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type request struct {
}

type response struct {
	Message string `json:"message"`
}

type client struct {
	address  string
	user     string
	password string
}

type Client interface {
	AccountCreate(name string, hostname string, origin string, stackname string) (*accountCreateResponse, error)
	// ApplicationCreate(hostname string, origin string, stackname string) (*CreatedApplication, error)
	// EnvironmentCreate(name string, sourceenvironmentname string, domainname string) (*CreatedEnvironment, error)
}

const (
	DefaultAddress = "https://aperture.section.io/api/v1"
)

func NewClient(user string, password string) (Client, error) {
	if user == "" {
		return nil, errors.New("user argument must not be empty.")
	}

	if password == "" {
		return nil, errors.New("password argument must not be empty.")
	}

	return &client{
		address:  DefaultAddress,
		user:     user,
		password: password,
	}, nil
}

func safeClose(c io.Closer, err *error) {
	if cerr := c.Close(); cerr != nil && *err == nil {
		*err = cerr
	}
}

func (c client) newRequest() request {
	return request{}
}

func (c client) httpPost(uri string, requestBody []byte) (responseBody []byte, outErr error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("POST", c.address+uri, bytes.NewReader(requestBody))

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.user, c.password)

	response, err := httpClient.Do(req)

	if err != nil {
		return nil, errors.Wrap(err, "HTTP POST failed for request.")
	}
	defer safeClose(response.Body, &outErr)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c client) httpPostJson(uri string, req interface{}, resp interface{}) error {

	reqBody, err := json.Marshal(req)
	if err != nil {
		return errors.Wrapf(err, "Failed to JSON encode request: %v", req)
	}

	respBody, err := c.httpPost(uri, reqBody)
	if err != nil {
		return errors.Wrapf(err, "Failed to HTTP POST request: %v", req)
	}

	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return errors.Wrapf(err, "Could not JSON decode response: %s", respBody)
	}

	return nil
}
