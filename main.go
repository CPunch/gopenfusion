package main

import (
	"log"

	"github.com/CPunch/gopenfusion/config"
	"github.com/CPunch/gopenfusion/db"
	"github.com/CPunch/gopenfusion/server"
)

func main() {
	dbHndlr, err := db.OpenLiteDB("test.db")
	if err != nil {
		log.Panicf("failed to open db: %v", err)
	}
	dbHndlr.Setup()

	loginServer, err := server.NewLoginServer(dbHndlr, config.LOGIN_PORT)
	if err != nil {
		log.Panicf("failed to create login server: %v", err)
	}

	shardServer, err := server.NewShardServer(dbHndlr, config.SHARD_PORT)
	if err != nil {
		log.Panicf("failed to create shard server: %v", err)
	}

	loginServer.AddShard(shardServer)
	go loginServer.Start()

	shardServer.Start()
}
