package main

import (
	"github.com/samaity/grpc-cache/server"
	"github.com/samaity/grpc-cache/client"
)

func main() {
	// Start the cache server in a goroutine
	go server.ServerMain()

	// Call the cache client
	client.ClientMain()

	// Keep the main goroutine running
	select {}
}
