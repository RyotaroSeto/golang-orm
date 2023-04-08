package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

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
	// var req createUserRequest

	// arg := db.CreateUserParams{
	// 	Username: req.Username,
	// 	FullName: req.FullName,
	// 	Email:    req.Email,
	// }

	// user, err := server.store.CreateUser(ctx, arg)
	// if err != nil {
	// 	if pqErr, ok := err.(*pq.Error); ok {
	// 		switch pqErr.Code.Name() {
	// 		case "unique_violation":
	// 			ctx.JSON(http.StatusForbidden, errorResponse(err))
	// 			return
	// 		}
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	rsp := newUserResponse()
	ctx.JSON(http.StatusOK, rsp)
}
