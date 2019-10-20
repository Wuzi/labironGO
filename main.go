package main

func main() {
	defer exit()

	game := game{}
	game.new(16, 110)
	game.run()
}
