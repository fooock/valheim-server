package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/valheim-server/valheim-agent/internal"
)

func main() {
	host := flag.String("host", "localhost", "Host to attach this service agent")
	port := flag.Int("port", 8080, "Port to bind this service agent")
	game := flag.String("game", "", "Address of the game server")

	flag.Parse()

	// check if game address is valid
	if *game == "" {
		log.Fatalln("Cannot start agent because game address is empty. Please set --game flag")
	}

	// create agent configuration and start it
	config := internal.AgentConfiguration{
		Host: *host,
		Port: *port,
		Game: *game,
	}
	agent := internal.Agent{Config: config}

	err := agent.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Agent started")

	// handle exit signal to shutdown the server
	exitChann := make(chan os.Signal, 1)
	signal.Notify(exitChann, os.Interrupt)
	<-exitChann

	err = agent.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Agent stopped")
}
