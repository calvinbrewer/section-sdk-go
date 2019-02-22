package api

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestEnvironmentCreate(t *testing.T) {
	defer gock.Off()

	const name = "test"
	const sourceEnvironmentName = "SomeEnv"
	const domainName = "www.test.com"

	responseBody := getTestDataString(t, "environment.create.responseBody.json")

	client := getTestClient(t)

	gock.New("https://aperture.section.io").
		Post("/api/v1/account/1234/application/4567/environment/create").
		Reply(200).
		BodyString(responseBody)

	resp, err := client.EnvironmentCreate(1234, 4567, name, sourceEnvironmentName, domainName)
	assert.Nil(t, err, "AccountCreate error")

	assert.Equal(t, 1234, resp.EnvironmentID, "EnvironmentID")

	assert.True(t, gock.IsDone(), "gock.IsDone")
}
