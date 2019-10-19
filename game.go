package main

import "math/rand"

type game struct {
	grid         [][]tile
	attemptsLeft int
	generating   bool
}

// new creates a new game
func (g *game) new(sizeX, sizeY int) {
	g.attemptsLeft = 5
	g.generating = true

	// initialize grid
	g.grid = make([][]tile, sizeX)
	for i := range g.grid {
		g.grid[i] = make([]tile, sizeY)
	}

	// fill the grid with walls
	for x := 0; x < sizeX; x++ {
		for y := 0; y < sizeY; y++ {
			g.grid[x][y] = wall
		}
	}

	// create a starting point in the middle of first row to generate from
	g.grid[0][len(g.grid[0])/2] = ground

	for g.generating {
		g.generate()
	}
}

// generates a new map
func (g *game) generate() {
	pCount := 0

	for x := 0; x < len(g.grid); x++ {
		for y := 0; y < len(g.grid[0]); y++ {
			if g.getTile(x, y) == ground {
				pCount += g.makePassage(x, y, -1, 0)
				pCount += g.makePassage(x, y, 1, 0)
				pCount += g.makePassage(x, y, 0, -1)
				pCount += g.makePassage(x, y, 0, 1)
			}
		}
	}

	if pCount == 0 {
		g.attemptsLeft--
		if g.attemptsLeft < 0 {
			possibleExits := []int{}
			for x := 0; x < len(g.grid); x++ {
				if g.getTile(x, len(g.grid[0])-2) == ground {
					possibleExits = append(possibleExits, x)
				}
			}

			// create a random exit
			x := possibleExits[rand.Intn(len(possibleExits))]
			g.setTile(x, len(g.grid[0])-1, ground)

			g.generating = false
		}
	}
}

// makePassage checks around a coordinate if it's all walls
// and randomly creates a passage
func (g *game) makePassage(x, y, i, j int) int {
	if g.getTile(x+i, y+j) == wall &&
		g.getTile(x+i+j, y+j+i) == wall &&
		g.getTile(x+i-j, y+j-i) == wall {
		if g.getTile(x+i+i, y+j+j) == wall &&
			g.getTile(x+i+i+j, y+j+j+i) == wall &&
			g.getTile(x+i+i-j, y+j+j-i) == wall {
			if rand.Float32() > 0.5 {
				g.setTile(x+i, y+j, ground)
				return 1
			}
		}
	}
	return 0
}

// getTile gets the type of a tile in a x and y coordinate
func (g *game) getTile(x, y int) tile {
	if x >= 0 && y >= 0 && x < len(g.grid) && y < len(g.grid[0]) {
		return g.grid[x][y]
	}
	return ground
}

// setTile sets the type of a tile of a x and y coordinate
func (g *game) setTile(x, y int, tile tile) {
	if x >= 0 && y >= 0 && x < len(g.grid) && y < len(g.grid[0]) {
		g.grid[x][y] = tile
	}
}
