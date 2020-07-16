package service

import (
	"testing"

	"github.com/sunil206b/go-microservices/repo"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFoundInDatabase(t *testing.T) {
	db := repo.GetUserByIDMock()
	defer db.Close()
	usersvc := NewUserService(db)
	user, err := usersvc.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User with id 0 not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	db := repo.GetUserByIDMock()
	defer db.Close()
	usersvc := NewUserService(db)
	user, err := usersvc.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Virat", user.FirstName)
	assert.EqualValues(t, "Kohli", user.LastName)
	assert.EqualValues(t, "virat@example.com", user.Email)
}
