package main

type player struct {
	x int
	y int
}

func (p *player) move(direction string, g *game) {
	x, y := p.x, p.y

	switch direction {
	case "UP":
		p.x--
		if p.x < 0 {
			p.x++
		}
	case "DOWN":
		p.x++
		if p.x == len(g.grid)-1 {
			p.x--
		}
	case "RIGHT":
		p.y++
		if p.y == len(g.grid[0]) {
			p.y--
		}
	case "LEFT":
		p.y--
		if p.y < 0 {
			p.y++
		}
	}

	if g.grid[p.x][p.y] == wall {
		p.x = x
		p.y = y
	}
}
