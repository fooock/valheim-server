package internal

import (
	"encoding/json"
	"net/http"

	"github.com/rumblefrog/go-a2s"
)

// GameService is used to interact with the gameserver
type GameService struct {
	Client *a2s.Client
}

type gameServer struct {
	Name             string `json:"name"`
	ConnectedPlayers int    `json:"connected_players"`
	MaxPlayers       int    `json:"max_players"`
}

// GetGameInfo return the gameserver info
func (g *GameService) GetGameInfo(w http.ResponseWriter, r *http.Request) {
	info, err := g.Client.QueryInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	gameServer := gameServer{
		Name:             info.Name,
		MaxPlayers:       int(info.MaxPlayers),
		ConnectedPlayers: int(info.Players),
	}
	result, err := json.Marshal(gameServer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
