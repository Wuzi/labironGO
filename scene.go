package main

import (
	"fmt"
	"time"
)

type scene int

const (
	intro    scene = 0
	gameplay scene = 1
	escaped  scene = 2
	gameover scene = 3
	quit     scene = 4
)

func (g *game) drawGameplayScene() {
	clearScreen()

	for i, tiles := range g.grid {
		for j, tile := range tiles {
			if tile == wall {
				fmt.Printf("\x1b[42m \x1b[0m")
			} else {
				fmt.Print(" ")
			}

			if j == len(tiles)-1 {
				fmt.Print("\n")
				moveCursor(i+1, 0)
			}
		}
	}

	moveCursor(g.player.x, g.player.y)
	fmt.Printf("\x1b[41m \x1b[0m")

	moveCursor(len(g.grid)+2, 0)
	fmt.Printf("Press ESC to give up")

	moveCursor(len(g.grid), 0)
}

func (g *game) drawIntroScene() {
	clearScreen()

	fmt.Println("You wake up in the middle of a maze...")
	time.Sleep(5 * time.Second)

	g.scene = gameplay
}

func (g *game) drawEscapedScene() {
	clearScreen()

	fmt.Println("Congratulations! You escaped the maze.")
	time.Sleep(5 * time.Second)

	clearScreen()
	g.scene = quit
}

func (g *game) drawGameoverScene() {
	clearScreen()

	fmt.Println("You died inside the maze")
	time.Sleep(5 * time.Second)

	clearScreen()
	g.scene = quit
}
