package main

func main() {
	snakes := map[int]int{16: 6, 47: 26, 49: 11, 56: 53, 62: 19, 64: 60, 87: 24, 93: 73, 95: 75, 98: 78}
	ladders := map[int]int{1: 38, 4: 14, 9: 31, 21: 42, 28: 84, 36: 44, 51: 67, 71: 91, 80: 100}
	snakeBoard := NewSnakeLadderBoard(100, snakes, ladders)

	dice := &SixSideDice{}

	players := []*Player{
		{Name: "Sristi", ID: "1"},
		{Name: "Rohit", ID: "2"},
	}

	game := &SnakeLadderBoardGame{
		Board:   snakeBoard,
		Dice:    dice,
		Players: players,
	}

	game.Start()
}
