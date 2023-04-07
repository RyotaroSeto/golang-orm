package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func newUserResponse() userResponse {
	return userResponse{
		Username: "test",
		FullName: "test",
		Email:    "test",
	}
}

func sqlcCreateUser(ctx *gin.Context) {
	rsp := newUserResponse()
	ctx.JSON(http.StatusOK, rsp)
}
