package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountGet(t *testing.T) {
	defer gock.Off()

	const accountID = 1469
	responseBody := getTestDataString(t, "account.get.responseBody.json")

	client := getTestClient(t)

	gock.New("https://aperture.section.io").
		Post("/api/v1/account/" + accountID).
		Reply(200).
		BodyString(responseBody)

	account, err := client.AccountGet(accountID)
	assert.Nil(t, err, "AccountGet error")

	assert.Equal(t, 1469, account.id, "accountID")

	assert.True(t, gock.IsDone(), "gock.IsDone")
}
