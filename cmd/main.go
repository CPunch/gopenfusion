package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/CPunch/gopenfusion/config"
	"github.com/CPunch/gopenfusion/internal/db"
	"github.com/CPunch/gopenfusion/internal/redis"

	"github.com/google/subcommands"
)

var dbHndlr *db.DBHandler
var redisHndlr *redis.RedisHandler

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&shardCommand{}, "")
	subcommands.Register(&loginCommand{}, "")

	var err error
	dbHndlr, err = db.OpenPostgresDB(config.GetDBAddr())
	if err != nil {
		log.Panicf("failed to open db: %v", err)
	}

	if err = dbHndlr.Setup(); err != nil {
		log.Panicf("failed to setup db: %v", err)
	}

	redisHndlr, err = redis.OpenRedis(config.GetRedisAddr())
	if err != nil {
		log.Panicf("failed to open redis: %v", err)
	}

	flag.Parse()
	os.Exit(int(subcommands.Execute(context.Background())))
}
