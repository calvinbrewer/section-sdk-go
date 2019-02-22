package api

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestAccountCreate(t *testing.T) {
	defer gock.Off()

	const name = "test"
	const hostname = "www.test.com"
	const origin = "origin.test.com"
	const stackname = "stack"

	responseBody := getTestDataString(t, "account.get.responseBody.json")

	client := getTestClient(t)

	gock.New("https://aperture.section.io").
		Post("/api/v1/account/create").
		Reply(200).
		BodyString(responseBody)

	accountResp, err := client.AccountCreate(name, hostname, origin, stackname)
	assert.Nil(t, err, "AccountCreate error")

	assert.Equal(t, 1469, accountResp.AccountID, "accountID")

	assert.True(t, gock.IsDone(), "gock.IsDone")
}
