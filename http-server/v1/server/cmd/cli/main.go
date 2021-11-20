package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/zshainsky/learning_go_with_testing/http-server/v1/server"
)

const dbFileName = "game.db.json"

func main() {

	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	defer close()
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	// game := poker.NewGame(poker.BlindAlertFunc(poker.StdOutAlerter), store)
	game := poker.NewTexasHoldem(poker.BlindAlertFunc(poker.StdOutAlerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdin, game)
	cli.PlayPoker()
}
