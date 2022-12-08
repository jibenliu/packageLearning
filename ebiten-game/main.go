package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

//go:generate go install github.com/hajimehoshi/file2byteslice
//go:generate mkdir -p resources
//go:generate file2byteslice -input ./images/ship.png -output resources/ship.go -package resources -var ShipPng
//go:generate file2byteslice -input ./images/alien.png -output resources/alien.go -package resources -var AlienPng
//go:generate file2byteslice -input config.json -output resources/config.go -package resources -var ConfigJson
func main() {
	ebiten.SetWindowTitle("外星人入侵")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
