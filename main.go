package main

import "github.com/CPunch/GopenFusion/server"

func main() {
	server := server.NewLoginServer()
	server.Start()
}
