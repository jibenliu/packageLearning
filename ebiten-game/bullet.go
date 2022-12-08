package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Bullet struct {
	GameObject
	image       *ebiten.Image
	speedFactor float64
}

func NewBullet(cfg *Config, ship *Ship) *Bullet {
	rect := image.Rect(0, 0, cfg.BulletWidth, cfg.BulletHeight)
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(cfg.BulletColor)

	b := &Bullet{}
	b.width = cfg.BulletWidth
	b.height = cfg.BulletHeight
	b.x = ship.x + float64(ship.width-cfg.BulletWidth)/2
	b.y = float64(cfg.ScreenHeight - ship.height - cfg.BulletHeight)
	b.speedFactor = cfg.BulletSpeedFactor
	b.image = img
	return b
}

func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.x, bullet.y)
	screen.DrawImage(bullet.image, op)
}

func (bullet *Bullet) outOfScreen() bool {
	return bullet.y < -float64(bullet.height)
}
