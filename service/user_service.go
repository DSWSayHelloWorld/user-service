package service

import (
	"database/sql"

	"github.com/sunil206b/go-microservices/model"
	"github.com/sunil206b/go-microservices/repo"
	"github.com/sunil206b/go-microservices/utils"
)

type UserService struct {
	repo repo.IUserRepo
}

func NewUserService(conn *sql.DB) *UserService {
	return &UserService{
		repo: repo.NewUserRepo(conn),
	}
}

func (u *UserService) GetUser(id uint64) (*model.User, *utils.ApplicationError) {
	return u.repo.GetUserByID(id)
}
