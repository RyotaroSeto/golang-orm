package api

import (
	"golang-orm/util"
	"log"

	"github.com/gin-gonic/gin"
)

func NewServer(config util.Config) {
	r := gin.Default()

	r.GET("/sqlc/create", func(ctx *gin.Context) {
		sqlcCreateUser(ctx, config)
	})

	r.GET("/sqlx/create", func(ctx *gin.Context) {
		sqlxCreateUser(ctx, config)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
