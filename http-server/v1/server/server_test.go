package poker_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	poker "github.com/zshainsky/learning_go_with_testing/http-server/v1/server"
)

func TestGETPlayers(t *testing.T) {
	store := poker.NewStupPlayerStore(
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	)
	server := poker.NewPlayerServer(store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})
}

func TestStoreWins(t *testing.T) {
	store := poker.NewStupPlayerStore(
		map[string]int{},
		nil,
		nil,
	)

	server := poker.NewPlayerServer(store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusAccepted)
		poker.AssertPlayerWin(t, store, "Pepper")

	})

}
func TestLeague(t *testing.T) {
	store := poker.StubPlayerStore{}
	server := poker.NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []poker.Player

		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}
		poker.AssertStatus(t, response, http.StatusOK)
	})
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []poker.Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := poker.NewStupPlayerStore(nil, nil, wantedLeague)
		server := poker.NewPlayerServer(store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := poker.GetLeagueFromResponse(t, response.Body)
		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertLeague(t, got, wantedLeague)
		poker.AssertContentType(t, response, "application/json")
	})
}

func TestGame(t *testing.T) {
	server := poker.NewPlayerServer(&poker.StubPlayerStore{})
	t.Run("GET /game returns 200", func(t *testing.T) {
		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
	})

}

func newGameRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}
func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}
