//+build wireinject

package main

import "github.com/google/wire"

func InitMission(name string) Mission {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}
}

// 上面的代码是wire命令需要的
// 在当前目录下执行wire命令,即可得到wire_gen.go命令，在生成的代码下添加main函数 也就可以实现base的业务逻辑

