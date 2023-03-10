package main

import (
	"github.com/CPunch/gopenfusion/db"
	"github.com/CPunch/gopenfusion/server"
)

func main() {
	db.DefaultDB, _ = db.OpenLiteDB("test.db")
	db.DefaultDB.Setup()

	server := server.NewLoginServer()
	server.Start()
}
