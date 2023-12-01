package main

import (
	"context"
	"flag"
	"log"

	"github.com/CPunch/gopenfusion/config"
	"github.com/CPunch/gopenfusion/shard"
	"github.com/google/subcommands"
)

type shardCommand struct {
	port int
}

func (s *shardCommand) Name() string {
	return "shard"
}

func (s *shardCommand) Synopsis() string {
	return "Starts shard service"
}

func (s *shardCommand) Usage() string {
	return s.Name() + " - " + s.Synopsis() + ":\n"
}

func (s *shardCommand) SetFlags(f *flag.FlagSet) {
	f.IntVar(&s.port, "port", config.SHARD_PORT, "Hosts the service on this port")
}

func (s *shardCommand) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	shardServer, err := shard.NewShardServer(ctx, dbHndlr, redisHndlr, s.port)
	if err != nil {
		log.Panicf("failed to create shard server: %v", err)
	}

	shardServer.Start()

	return subcommands.ExitSuccess
}
