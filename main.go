package main

import (
	"github.com/MarkTBSS/080_Pagination/config"
	"github.com/MarkTBSS/080_Pagination/databases"
	"github.com/MarkTBSS/080_Pagination/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
