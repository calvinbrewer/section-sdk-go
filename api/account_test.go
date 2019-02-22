package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func TestAccountGet(t *testing.T) {
	defer gock.Off()

	const accountID = 1469
	responseBody := getTestDataString(t, "account.get.responseBody.json")

	client := getTestClient(t)

	gock.New("https://aperture.section.io").
		Post("/api/v1/account/create").
		Reply(200).
		BodyString(responseBody)

	account, err := client.AccountGet(accountID)
	assert.Nil(t, err, "AccountGet error")

	assert.Equal(t, 1469, account.AccountID, "accountID")

	assert.True(t, gock.IsDone(), "gock.IsDone")
}
