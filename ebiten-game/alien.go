package main

import (
	"bytes"
	"ebiten-game/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

type Alien struct {
	GameObject
	image       *ebiten.Image
	speedFactor float64
}

func NewAlien(cfg *Config) *Alien {
	//img, _, err := ebitenutil.NewImageFromFile("./images/alien.png")
	//if err != nil {
	//	log.Fatal(err)
	//}

	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(resources.AlienPng))
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	a := &Alien{}
	a.image = img
	a.width = width
	a.height = height
	a.x = 0
	a.y = 0
	a.speedFactor = cfg.AlienSpeedFactor
	return a
}

func (alien *Alien) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(alien.x, alien.y)
	screen.DrawImage(alien.image, op)
}

func (alien *Alien) outOfScreen(cfg *Config) bool {
	return alien.y < -float64(alien.height)
}
