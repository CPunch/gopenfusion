package main

import (
	"github.com/CPunch/GopenFusion/db"
	"github.com/CPunch/GopenFusion/server"
)

func main() {
	db.DefaultDB, _ = db.OpenLiteDB("test.db")
	db.DefaultDB.Setup()

	server := server.NewLoginServer()
	server.Start()
}
