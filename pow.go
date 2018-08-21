package main

// import (
// 	"crypto/sha256"
// 	"encoding/hex"
// 	"fmt"
// 	"github.com/davecgh/go-spew/spew"
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// 	"strings"
// 	"sync"
// 	"time"
// )

// /*
// 	Index 是这个块在整个链中的位置
// 	Timestamp 显而易见就是块生成时的时间戳
// 	Hash 是这个块通过 SHA256 算法生成的散列值
// 	PrevHash 代表前一个块的 SHA256 散列值
// 	BPM 每分钟心跳数，也就是心率。还记得文章开头说到的吗？
// */
// type Block struct {
// 	Index      int
// 	Timestamp  string
// 	BPM        int
// 	Hash       string
// 	PreHash    string
// 	Difficulty int
// 	Nonce      string
// }

// var mutex = &sync.Mutex{}

// // 难度
// const difficulty = 3

// // 定义一个结构表示整个链，最简单的表示形式就是一个 Block 的 slice
// var Blockchain []Block

// // 计算散列值
// func calculateHash(block Block) string {
// 	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PreHash + block.Nonce
// 	h := sha256.New()
// 	h.Write([]byte(record))
// 	hashed := h.Sum(nil)
// 	return hex.EncodeToString(hashed)
// }

// // 生成区块
// func generateBlock(oldBlock Block, BPM int) Block {
// 	var newBlock Block
// 	t := time.Now()
// 	newBlock.Timestamp = t.String()
// 	newBlock.Index = oldBlock.Index + 1
// 	newBlock.BPM = BPM
// 	newBlock.PreHash = oldBlock.Hash
// 	newBlock.Difficulty = difficulty

// 	for i := 0; ; i++ {
// 		hex := fmt.Sprintf("%x", i)
// 		newBlock.Nonce = hex

// 		if isHashValid(calculateHash(newBlock), newBlock.Difficulty) {
// 			newBlock.Hash = calculateHash(newBlock)
// 			fmt.Println(calculateHash(newBlock), " work done!")
// 			break
// 		} else {
// 			fmt.Println(calculateHash(newBlock), " do more work!")
// 			time.Sleep(time.Millisecond)
// 			continue
// 		}
// 	}

// 	return newBlock
// }

// func isHashValid(hash string, difficulty int) bool {
// 	prefix := strings.Repeat("0", difficulty)
// 	return strings.HasPrefix(hash, prefix)
// }

// func isBlockValid(newBlock, oldBlock Block) bool {
// 	if oldBlock.Index+1 != newBlock.Index {
// 		return false
// 	}

// 	if oldBlock.Hash != newBlock.PreHash {
// 		return false
// 	}

// 	if calculateHash(newBlock) != newBlock.Hash {
// 		return false
// 	}

// 	return true
// }

// func replaceChain(newBlocks []Block) {
// 	if len(newBlocks) > len(Blockchain) {
// 		Blockchain = newBlocks
// 	}
// }

// func main() {
// 	router := gin.Default()

// 	go func() {
// 		genesisBlock := Block{0, time.Now().String(), 0, "", "", difficulty, ""}

// 		mutex.Lock()
// 		Blockchain = append(Blockchain, genesisBlock)
// 		mutex.Unlock()
// 	}()

// 	router.GET("/", handleGetBlockChain)
// 	router.POST("/", handleWriterBlockChan)

// 	router.Run(":5000")
// }

// func handleGetBlockChain(c *gin.Context) {
// 	c.JSON(200, Blockchain)
// }

// func handleWriterBlockChan(c *gin.Context) {
// 	var bpm struct{ BPM int }
// 	if err := c.BindJSON(&bpm); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	//ensure atomicity when creating new block
// 	mutex.Lock()
// 	oldBlock := Blockchain[len(Blockchain)-1]
// 	newBlock := generateBlock(oldBlock, bpm.BPM)
// 	mutex.Unlock()

// 	if isBlockValid(newBlock, oldBlock) {
// 		newBlockChain := append(Blockchain, newBlock)
// 		replaceChain(newBlockChain)
// 	}

// 	spew.Dump(newBlock)

// 	c.JSON(http.StatusCreated, newBlock)
// }
