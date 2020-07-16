package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/sunil206b/go-microservices/service"
	"github.com/sunil206b/go-microservices/utils"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		svc: service.NewUserService(db),
	}
}

// GetUser function will handle the user request
func (u *UserHandler) GetUser(c *gin.Context) {
	userIDParam := c.Param("id")
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	var errMsg *utils.ApplicationError
	if err != nil {
		errMsg.Message = "Not a valid user id"
		errMsg.Code = "bad_request"
		utils.ResponseError(c, http.StatusBadRequest, errMsg)
		return
	}

	user, errMsg := u.svc.GetUser(userID)
	if errMsg != nil {
		utils.ResponseError(c, http.StatusNotFound, errMsg)
		return
	}
	utils.Response(c, http.StatusOK, user)
}
