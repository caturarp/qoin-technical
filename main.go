package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/caturarp/qoin-technical.git/game"
)

func main() {
	player := flag.Int("player", 3, "Player Count")
	dice := flag.Int("dice", 4, "Dice Count")

	flag.Parse()
	winner, err := game.GameOfDice(*player, *dice)
	if err != nil {
		log.Print(err.Error())
		return
	}
	fmt.Printf("Winner: %v\n", winner)
}
