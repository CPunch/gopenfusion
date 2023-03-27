package main

import (
	"log"

	"github.com/CPunch/gopenfusion/config"
	"github.com/CPunch/gopenfusion/core/db"
	"github.com/CPunch/gopenfusion/login"
	"github.com/CPunch/gopenfusion/shard"
)

func main() {
	dbHndlr, err := db.OpenLiteDB("test.db")
	if err != nil {
		log.Panicf("failed to open db: %v", err)
	}
	dbHndlr.Setup()

	loginServer, err := login.NewLoginServer(dbHndlr, config.LOGIN_PORT)
	if err != nil {
		log.Panicf("failed to create login server: %v", err)
	}

	shardServer, err := shard.NewShardServer(dbHndlr, config.SHARD_PORT)
	if err != nil {
		log.Panicf("failed to create shard server: %v", err)
	}

	loginServer.AddShard(shardServer)
	go loginServer.Start()

	shardServer.Start()
}
