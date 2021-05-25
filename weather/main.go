package main

import "weather/cmd"

func main() {
	_ = cmd.Execute()
}

// ./weather name -n=武汉
// ./weather code -c=430000