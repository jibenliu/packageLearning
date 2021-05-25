package core

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index         int64  //区块编号
	Timestamp     int64  //区块时间戳
	PrevBlockHash string //上一个区块的hash值
	Hash          string //上一个区块的hash值

	Data string //区块数据
}

func calculateHash(b Block) string {
	blockData := strconv.FormatInt(b.Index, 10) + strconv.FormatInt(b.Timestamp, 10) + b.PrevBlockHash + b.Data
	hashBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashBytes[:])
}

func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}
