package main

import (
	"bytes"
	"ebiten-game/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

type Ship struct {
	GameObject
	image *ebiten.Image
}

func NewShip(screenWidth, screenHeight int) *Ship {
	//img, _, err := ebitenutil.NewImageFromFile("./images/ship.png")
	//if err != nil {
	//	log.Fatal(err)
	//}

	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(resources.ShipPng))
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	ship := &Ship{}
	ship.x = float64(screenWidth-width) / 2
	ship.y = float64(screenHeight-height) / 2
	ship.width = width
	ship.height = height
	ship.image = img

	return ship
}

func (ship *Ship) Draw(screen *ebiten.Image, cfg *Config) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cfg.ScreenWidth-ship.width)/2, float64(cfg.ScreenHeight-ship.height))
	screen.DrawImage(ship.image, op)
}
