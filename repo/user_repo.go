package repo

import (
	"fmt"

	"github.com/sunil206b/go-microservices/model"
	"github.com/sunil206b/go-microservices/utils"
)

var (
	users = map[uint64]*model.User{
		123: {ID: 123, FirstName: "Virat", LastName: "Kohli", Email: "virat@example.com"},
	}
)

// GetUser function will return
func GetUser(id uint64) (*model.User, *utils.ApplicationError) {
	user := users[id]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message: fmt.Sprintf("User with id %d not found", id),
			Code:    "not_found",
		}
	}
	return user, nil
}
