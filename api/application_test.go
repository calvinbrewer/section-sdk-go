package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func TestApplicationCreate(t *testing.T) {
	defer gock.Off()

	const accountID = 1
	const hostname = "www.test.com"
	const origin = "origin.test.com"
	const stackname = "stack"

	responseBody := getTestDataString(t, "application.create.responseBody.json")

	client := getTestClient(t)

	gock.New("https://aperture.section.io").
		Post("/api/v1/account/1/application/create").
		Reply(200).
		BodyString(responseBody)

	applicationResp, err := client.ApplicationCreate(accountID, hostname, origin, stackname)
	assert.Nil(t, err, "ApplicationCreate error")

	assert.Equal(t, 1, applicationResp.ApplicationID, "applicationID")

	assert.True(t, gock.IsDone(), "gock.IsDone")
}
