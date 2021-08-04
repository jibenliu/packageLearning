//+build wireinject

package main

import "github.com/google/wire"

func InitMission(p PlayerParam, m MonsterParam) (Mission, error) {
	wire.Build(NewPlayer, NewMonster, NewMission)
	return Mission{}, nil
}
