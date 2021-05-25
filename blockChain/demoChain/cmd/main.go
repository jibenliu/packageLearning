package main

import "newTest/blockChain/demoChain/core"

func main()  {
	bc := core.NewBlockChain() //创世区块
	bc.SendData("Send 1 BTC to Jacky")
	bc.SendData("Send 1 BTC to Jack")
	bc.PrintBlocks()
}
