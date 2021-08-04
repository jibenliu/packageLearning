package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

//区块链块的结构体
type Block struct {
	//Index 是这个块在整个链中的位置
	Index int
	//Timestamp 块生成时的时间戳
	Timestamp string
	//每分钟心跳数，也就是心率
	BPM int
	//通过 SHA256 算法生成的散列值
	Hash string
	//代表前一个块的 SHA256 散列值
	PrevHash string
}

//表示整个链
var blockChain []Block

//计算数据的sha256散列值
func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//生成块
func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

//校验块
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

//将过期的链更新为最新的链
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(blockChain) {
		blockChain = newBlocks
	}
}

//初始化web服务
func run() error {
	muxRouter := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on", os.Getenv("ADDR"))
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        muxRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockChain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

//get请求的handler
func handleGetBlockChain(w http.ResponseWriter, _ *http.Request) {
	bytes, err := json.MarshalIndent(blockChain, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = io.WriteString(w, string(bytes))
}

//post请求的payload
type Message struct {
	BPM int
}

//post请求的handler
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWidthJSON(w, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	newBlock, err := generateBlock(blockChain[len(blockChain)-1], m.BPM)
	if err != nil {
		respondWidthJSON(w, http.StatusCreated, newBlock)
	}


	if isBlockValid(newBlock, blockChain[len(blockChain)-1]) {
		newBlockChain := append(blockChain, newBlock)
		replaceChain(newBlockChain)
		spew.Dump(blockChain)
	}

	respondWidthJSON(w, http.StatusCreated, newBlock)
}

//返回客户端响应
func respondWidthJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("HTTP 500: Internal Server Error"))
	}
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//初始化区块链
	go func() {
		t := time.Now()
		genesisBlock := Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		blockChain = append(blockChain, genesisBlock)
	}()
	log.Fatal(run())
}
