package main

import "fmt"

type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{Name: "kitty"}
}

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{
		Name: "name",
	}
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player, m.Monster)
}

type EndingA struct {
	Player  Player
	Monster Monster
}

func NewEndingA(p Player, m Monster) EndingA {
	return EndingA{p, m}
}

func (p EndingA) Appear() {
	fmt.Printf("%s defeats %s, world peace!\n", p.Player.Name, p.Monster.Name)
}

type EndingB struct {
	Player  Player
	Monster Monster
}

func NewEndingB(p Player, m Monster) EndingB {
	return EndingB{p, m}
}

func (p EndingB) Appear() {
	fmt.Printf("%s defeats %s, but become monster, world darker!\n", p.Player.Name, p.Monster.Name)
}
