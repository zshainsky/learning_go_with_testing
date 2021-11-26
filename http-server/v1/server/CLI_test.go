package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	poker "github.com/zshainsky/learning_go_with_testing/http-server/v1/server"
)

type GameSpy struct {
	StartedWith int
	StartCalled bool

	BlindAlert     []byte
	FinishedWith   string
	FinishedCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int, out io.Writer) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
	out.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedCalled = true
	g.FinishedWith = winner
}

func TestCLI(t *testing.T) {
	var dummyStdOut = &bytes.Buffer{}

	t.Run("check 'Chris Wins' text using string as input to io.Reader", func(t *testing.T) {
		in := strings.NewReader("7\nChris wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		if game.FinishedWith != "Chris" {
			t.Errorf("expected finish called with 'Chris' but got %q", game.FinishedWith)
		}
	})
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("hello World")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)

	})
	t.Run("start game with 3 players and finish game with Chris as the winner", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("3", "Chris wins")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")

	})
	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("3", "Lloyd is a killer")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotFinished(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputMessage)

	})

}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
func assertGameNotStarted(t testing.TB, game *GameSpy) {
	t.Helper()

	if game.StartCalled {
		t.Errorf("game shoud not have started")
	}
}

func assertGameStartedWith(t testing.TB, game *GameSpy, got int) {
	t.Helper()

	if game.StartedWith != got {
		t.Errorf("wanted a game with %d players but got %d", got, game.StartedWith)
	}

}
func assertGameNotFinished(t testing.TB, game *GameSpy) {
	t.Helper()

	if game.FinishedCalled {
		t.Errorf("game shoud not have finished")
	}
}
func assertFinishCalledWith(t testing.TB, game *GameSpy, got string) {
	if game.FinishedWith != got {
		t.Errorf("expected finish called with %q but got %q", got, game.FinishedWith)
	}
}
func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}
