package main

import (
	"encoding/json"
	"io"
	"net/http"
	"/blockChain/demoChain/core"
)

var blockChain *core.BlockChain

func run() {
	http.HandleFunc("/blockchain/get", blockChainGetHandler)
	http.HandleFunc("/blockchain/write", blockChainWriteHandler)
	http.ListenAndServe("localhost:8888", nil)
}

func blockChainGetHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(blockChain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = io.WriteString(w, string(bytes))
}

func blockChainWriteHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockChain.SendData(blockData)
	blockChainGetHandler(w, r)
}

func main() {
	blockChain = core.NewBlockChain()
	run()
}
