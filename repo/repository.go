package repo

import (
	"github.com/sunil206b/go-microservices/model"
	"github.com/sunil206b/go-microservices/utils"
)

type IUserRepo interface {
	GetUserByID(id uint64) (*model.User, *utils.ApplicationError)
}
