package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rumblefrog/go-a2s"
)

// AgentConfiguration needed to start the agent
type AgentConfiguration struct {
	Host string
	Port int
	Game string
}

// Agent is the service that will manage the game server resources from
// an external point of view
type Agent struct {
	Config  AgentConfiguration
	server  http.Server
	service GameService
}

// Init the game server agent
func (a *Agent) Init() error {
	log.Printf("Start agent on %s:%d", a.Config.Host, a.Config.Port)

	// create client to query game server
	client, err := a2s.NewClient(a.Config.Game)
	if err != nil {
		return err
	}
	a.service = GameService{Client: client}

	router := mux.NewRouter()
	v1group := router.PathPrefix("/v1").Subrouter()
	v1group.HandleFunc("/info", a.service.GetGameInfo).Methods("GET")

	// create server to handle requests
	a.server = http.Server{
		Addr:    fmt.Sprintf("%s:%d", a.Config.Host, a.Config.Port),
		Handler: router,
	}
	go a.server.ListenAndServe()
	return nil
}

// Shutdown the game server agent
func (a *Agent) Shutdown() error {
	log.Println("Shutdown agent")

	err := a.server.Close()
	if err != nil {
		return err
	}
	err = a.service.Client.Close()
	return err
}
