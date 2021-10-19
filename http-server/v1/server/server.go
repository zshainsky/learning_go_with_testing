package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const jsonContentType = "application/json"

type PlayerStore interface {
	GetPlayerScore(string) int
	RecordWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router
	return p
}

func (p *PlayerServer) leagueHandler(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("content-type", jsonContentType)
	json.NewEncoder(writer).Encode(p.store.GetLeague())
	writer.WriteHeader(http.StatusOK)

}

func (p *PlayerServer) playersHandler(writer http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "/players/")

	switch request.Method {
	case http.MethodPost:
		p.processWin(writer, player)
	case http.MethodGet:
		p.showScore(writer, player)
	}
}

func (p *PlayerServer) showScore(writer http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(writer, fmt.Sprint(score))
}

func (p *PlayerServer) processWin(writer http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	writer.WriteHeader(http.StatusAccepted)
}

func GetPlayerScore(name string) int {
	if name == "Pepper" {
		return 20
	}

	if name == "Floyd" {
		return 10
	}

	return 0
}
