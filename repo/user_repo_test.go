package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization:

	//Execution:
	user, err := GetUser(0)

	// Validation:
	assert.Nil(t, user, "We are not expecting a user with id 0")
	assert.NotNil(t, err, "We are not expecting a user with id 0")
	assert.EqualValues(t, "not_found", err.Code, "We are expecting user not found error")
	assert.EqualValues(t, "User with id 0 not found", err.Message, "This is the not the error message if user not found")
	/*if user != nil {
		t.Error("We are not expecting a user with id 0")
	}

	if err == nil {
		t.Error("We are expecting an error when user id is 0")
	}

	if err.Code != "not_found" {
		t.Error("We are expecting user not found error")
	}*/
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err, "We are not expecting error here.")
	assert.NotNil(t, user, "We are expecting user here")
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Virat", user.FirstName)
	assert.EqualValues(t, "Kohli", user.LastName)
	assert.EqualValues(t, "virat@example.com", user.Email)
}
