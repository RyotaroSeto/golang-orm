package api

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func NewServer(db *sql.DB) {
	r := gin.Default()

	r.GET("/sqlc/create", sqlcCreateUser)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
