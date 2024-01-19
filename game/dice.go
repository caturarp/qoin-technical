package game

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	point       int
	dice        int
	diceNumbers []int
}

func GameOfDice(totalPlayers int, totalDice int) (winner int, err error) {
	err = errors.New("no winner")
	if totalPlayers < 1 || totalDice < 1 {
		return winner, errors.New("input invalid, player and dice must be minimum 1")
	}
	fmt.Printf("Pemain = %v, Dadu = %v\n", totalPlayers, totalDice)
	playersPoint := make(map[int]Player, totalPlayers)

	x := 1
	for x <= totalPlayers {
		playersPoint[x] = Player{point: 0, dice: totalDice}
		x += 1
	}
	for {
		for i := 1; i <= totalPlayers; i++ {
			if playersPoint[i].dice == 0 {
				continue
			}

			rand.Seed(time.Now().UnixNano())
			randomNumber := rand.Intn(6) + 1

			currentPlayer := playersPoint[i]
			currentPlayer.diceNumbers = append(currentPlayer.diceNumbers, randomNumber)
			playersPoint[i] = currentPlayer

			if randomNumber == 6 {
				currentPlayer.dice -= 1
				currentPlayer.point += 1
				playersPoint[i] = currentPlayer
			} else if randomNumber == 1 && i < totalPlayers {
				currentPlayer.dice -= 1
				passedPlayer := playersPoint[i+1]
				passedPlayer.dice += 1
				playersPoint[i+1] = passedPlayer
				playersPoint[i] = currentPlayer
			} else if randomNumber == 1 && i == totalPlayers {
				currentPlayer.dice -= 1
				passedPlayer := playersPoint[1]
				passedPlayer.dice += 1
				playersPoint[1] = passedPlayer
				playersPoint[i] = currentPlayer
			}

			fmt.Printf("Player %d: %+v\n", i, playersPoint[i])
		}

		for player, playerDetail := range playersPoint {
			if playerDetail.point >= totalDice {
				return player, nil
			}
		}
	}
}
