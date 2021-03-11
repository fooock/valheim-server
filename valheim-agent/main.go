package main

import (
	"flag"
	"log"
)

// default port to query Valheim info using the source
// server queries
const steamGameQueryPort = 2457

var (
	host = flag.String("--host", "localhost", "Host to attach this service agent")
	port = flag.Int("--port", 8080, "Port to bind this service agent")
)

func main() {
	flag.Parse()
	log.Printf("Start Valheim agent into %s:%d", *host, *port)
}
