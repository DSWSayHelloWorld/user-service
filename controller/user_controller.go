package controller

import (
	"net/http"
	"strconv"

	"github.com/sunil206b/go-microservices/service"
	"github.com/sunil206b/go-microservices/utils"
)

// GetUser function will handle the user request
func GetUser(w http.ResponseWriter, r *http.Request) {
	userIDParam := r.URL.Query().Get("id")
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	var errMsg *utils.ApplicationError
	if err != nil {
		errMsg.Message = "Not a valid user id"
		errMsg.Code = "not_valid"
		utils.SendError(w, http.StatusBadRequest, *errMsg)
		return
	}

	user, errMsg := service.GetUser(userID)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, *errMsg)
		return
	}
	utils.SendSuccess(w, http.StatusOK, user)
}
