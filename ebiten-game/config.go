package main

import (
	"bytes"
	"ebiten-game/resources"
	"encoding/json"
	"image/color"
	"log"
)

type Config struct {
	ScreenWidth       int        `json:"screenWidth"`
	ScreenHeight      int        `json:"screenHeight"`
	Title             string     `json:"title"`
	BgColor           color.RGBA `json:"bgColor"`
	ShipSpeedFactor   float64    `json:"shipSpeedFactor"`
	BulletWidth       int        `json:"bulletWidth"`
	BulletHeight      int        `json:"bulletHeight"`
	BulletColor       color.RGBA `json:"bulletColor"`
	MaxBulletNum      int        `json:"maxBulletNum"`
	BulletInterval    int64      `json:"bulletInterval"`
	BulletSpeedFactor float64    `json:"bulletSpeedFactor"`
	AlienSpeedFactor  float64    `json:"alienSpeedFactor"`
	TitleFontSize     int        `json:"titleFontSize"`
	FontSize          int        `json:"fontSize"`
	SmallFontSize     int        `json:"smallFontSize"`
}

func loadConfig() *Config {
	//f, err := os.Open("./config.json")
	//if err != nil {
	//	log.Fatalf("os.Open failed: %v\n", err)
	//}
	//
	//
	//
	//var cfg Config
	//err = json.NewDecoder(f).Decode(&cfg)
	//if err != nil {
	//	log.Fatalf("json.Decode failed: %v\n", err)
	//}

	var cfg Config
	if err := json.NewDecoder(bytes.NewReader(resources.ConfigJson)).Decode(&cfg); err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}

	return &cfg
}
