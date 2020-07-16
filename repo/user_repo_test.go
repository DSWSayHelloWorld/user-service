package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization:
	db := GetUserByIDMock()
	defer db.Close()
	usrRepo := NewUserRepo(db)

	//Execution:
	user, errMsg := usrRepo.GetUserByID(0)

	// Validation:
	assert.Nil(t, user, "We are not expecting a user with id 0")
	assert.NotNil(t, errMsg, "We are not expecting a user with id 0")
	assert.EqualValues(t, "not_found", errMsg.Code, "We are expecting user not found error")
	assert.EqualValues(t, "User with id 0 not found", errMsg.Message, "This is the not the error message if user not found")
}

func TestGetUserNoError(t *testing.T) {
	db := GetUserByIDMock()
	defer db.Close()
	usrRepo := NewUserRepo(db)
	user, errMsg := usrRepo.GetUserByID(123)

	assert.Nil(t, errMsg, "We are not expecting error here.")
	assert.NotNil(t, user, "We are expecting user here")
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Virat", user.FirstName)
	assert.EqualValues(t, "Kohli", user.LastName)
	assert.EqualValues(t, "virat@example.com", user.Email)
}
