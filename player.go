package main

type Player struct {
	Name string
	ID   string
	Pos  int
}

func (p *Player) Position() int {
	return p.Pos
}

func (p *Player) Move(steps int) {
	p.Pos += steps
}
