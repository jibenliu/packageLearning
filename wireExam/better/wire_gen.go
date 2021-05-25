// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitMission(p PlayerParam, m MonsterParam) (Mission, error) {
	player, err := NewPlayer(p)
	if err != nil {
		return Mission{}, err
	}
	monster := NewMonster(m)
	mission := NewMission(player, monster)
	return mission, nil
}