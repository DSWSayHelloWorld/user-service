package service

import (
	"github.com/sunil206b/go-microservices/model"
	"github.com/sunil206b/go-microservices/repo"
	"github.com/sunil206b/go-microservices/utils"
)

func GetUser(id uint64) (*model.User, *utils.ApplicationError) {
	return repo.GetUser(id)
}
