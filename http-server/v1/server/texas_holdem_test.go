package poker_test

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	poker "github.com/zshainsky/learning_go_with_testing/http-server/v1/server"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		dummyPlayerStore := &poker.StubPlayerStore{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(5, ioutil.Discard)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100, To: ioutil.Discard},
			{At: 10 * time.Minute, Amount: 200, To: ioutil.Discard},
			{At: 20 * time.Minute, Amount: 300, To: ioutil.Discard},
			{At: 30 * time.Minute, Amount: 400, To: ioutil.Discard},
			{At: 40 * time.Minute, Amount: 500, To: ioutil.Discard},
			{At: 50 * time.Minute, Amount: 600, To: ioutil.Discard},
			{At: 60 * time.Minute, Amount: 800, To: ioutil.Discard},
			{At: 70 * time.Minute, Amount: 1000, To: ioutil.Discard},
			{At: 80 * time.Minute, Amount: 2000, To: ioutil.Discard},
			{At: 90 * time.Minute, Amount: 4000, To: ioutil.Discard},
			{At: 100 * time.Minute, Amount: 8000, To: ioutil.Discard},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		dummyPlayerStore := &poker.StubPlayerStore{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(7, ioutil.Discard)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100, To: ioutil.Discard},
			{At: 12 * time.Minute, Amount: 200, To: ioutil.Discard},
			{At: 24 * time.Minute, Amount: 300, To: ioutil.Discard},
			{At: 36 * time.Minute, Amount: 400, To: ioutil.Discard},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	dummyBlindAlerter := &poker.SpyBlindAlerter{}
	game := poker.NewTexasHoldem(dummyBlindAlerter, store)
	winner := "Ruth"

	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(cases []poker.ScheduledAlert, t *testing.T, blindAlerter *poker.SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {

			if len(blindAlerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
			}

			got := blindAlerter.Alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got +%v, want +%v", got, want)
	}
}
