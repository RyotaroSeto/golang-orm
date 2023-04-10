package main

import (
	"golang-orm/api"
	"golang-orm/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	api.NewServer(config)
}
