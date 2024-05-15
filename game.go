package main

import (
	"fmt"
	"time"
)

type Game interface {
	Start()
}

type SnakeLadderBoardGame struct {
	Board   Board
	Players []*Player
	Dice    Dice
}

func (g *SnakeLadderBoardGame) Start() {
	curr := 0
	for {
		player := g.Players[curr]
		points := g.Dice.Roll()
		fmt.Printf("Dice roll got %s %d points\n", player.Name, points)

		newPos := player.Position() + points
		if newPos > g.Board.Size() {
			newPos = player.Position()
		} else {
			newPos = g.Board.GetBoardPosition(newPos)
		}

		steps := newPos - player.Position()
		player.Move(steps)
		fmt.Printf("Player %s moved to %d\n", player.Name, player.Position())

		if player.Position() == g.Board.Size() {
			fmt.Printf("Player %s has won the game!!", player.Name)
			break
		}

		curr = (curr + 1) % (len(g.Players))
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}
