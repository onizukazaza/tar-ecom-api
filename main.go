package main

import (
	"github.com/onizukazaza/tar-ecom-api/config"
	"github.com/onizukazaza/tar-ecom-api/databases"
	"github.com/onizukazaza/tar-ecom-api/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewFiberServer(conf , db.Connect())

	server.Start()

	
}