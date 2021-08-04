package main

import (
	"learnGo/factoryMode/base"
	_ "learnGo/factoryMode/cls1" //匿名引用cls1包，自动注册
	_ "learnGo/factoryMode/cls2" //匿名引用cls2包，自动注册
)

func main() {
	c1 := base.Create("Class1")
	c1.Do()

	c2 := base.Create("Class2")
	c2.Do()
}
