package main

import "fmt"

type Board interface {
	Size() int
	GetBoardPosition(pos int) int
}

type SnakeLadderBoard struct {
	snakes  map[int]int
	ladders map[int]int
	size    int
}

func NewSnakeLadderBoard(size int, snakes map[int]int, ladders map[int]int) Board {
	return &SnakeLadderBoard{
		size:    size,
		snakes:  snakes,
		ladders: snakes,
	}
}

func (b *SnakeLadderBoard) GetBoardPosition(pos int) int {
	if end, ok := b.snakes[pos]; ok {
		fmt.Printf("Position %d got bitten by snake and moved to %d\n", pos, end)
		return end
	}

	if end, ok := b.ladders[pos]; ok {
		fmt.Printf("Position %d got ladder and moved to %d\n", pos, end)
		return end
	}

	return pos
}

func (b *SnakeLadderBoard) Size() int {
	return b.size
}
