package api

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestClient(t *testing.T) Client {
	const user = "dummy-user"
	const password = "dummy-password"
	client, err := NewClient(user, password)
	assert.Nil(t, err, "NewClient error")
	return client
}

func getTestDataString(t *testing.T, filename string) string {
	bytes, err := ioutil.ReadFile(path.Join("testdata", filename))
	assert.Nil(t, err, "ReadFile(%s) error", filename)
	return string(bytes)
}
