package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func calculateHash(toBeHashed string) string {
	hashBytes := sha256.Sum256([]byte(toBeHashed))
	hashInStr := hex.EncodeToString(hashBytes[:])
	log.Printf("%s, %s", toBeHashed, hashInStr)
	return hashInStr
}

func main() {
	calculateHash("test1")
}
