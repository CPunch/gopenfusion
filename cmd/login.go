package main

import (
	"context"
	"flag"
	"log"

	"github.com/CPunch/gopenfusion/internal/config"
	"github.com/CPunch/gopenfusion/login"
	"github.com/google/subcommands"
)

type loginCommand struct {
	port int
}

func (s *loginCommand) Name() string {
	return "login"
}

func (s *loginCommand) Synopsis() string {
	return "Starts login service"
}

func (s *loginCommand) Usage() string {
	return s.Name() + " - " + s.Synopsis() + ":\n"
}

func (s *loginCommand) SetFlags(f *flag.FlagSet) {
	f.IntVar(&s.port, "port", config.LOGIN_PORT, "Hosts the service on this port")
}

func (s *loginCommand) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	loginServer, err := login.NewLoginServer(ctx, dbHndlr, redisHndlr, s.port)
	if err != nil {
		log.Panicf("failed to create shard server: %v", err)
	}

	loginServer.Start()

	return subcommands.ExitSuccess
}
