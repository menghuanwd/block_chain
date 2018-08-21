package main

// import (
// 	"block_chain/merkletree"
// 	"block_chain/models"
// 	"block_chain/services"
// 	"block_chain/utils"
// 	"fmt"
// 	"github.com/davecgh/go-spew/spew"
// 	"github.com/gin-gonic/gin"
// 	"math/rand"
// 	"net/http"
// 	"strconv"
// 	"sync"
// 	"time"
// )

// var mutex = &sync.Mutex{}

// // 节点的存储map，同时也会保存每个节点持有的令牌数
// var validators = make(map[string]int)

// // 定义一个结构表示整个链，最简单的表示形式就是一个 Block 的 slice
// var Blockchain []models.Block

// // 临时存储单元，在区块被选出来并添加到BlockChain之前，临时存储在这里
// var tempBlocks []models.Block

// // Block的通道，任何一个节点在提出一个新块时都将它发送到这个通道
// var candidateBlocks = make(chan models.Block)

// func GenTempBlocks() {
// 	for candidateBlock := range candidateBlocks {
// 		tempBlocks = append(tempBlocks, candidateBlock)
// 	}
// }

// func HandleGetBlockChain(c *gin.Context) {
// 	c.JSON(200, Blockchain)
// }

// func HandlePay(c *gin.Context) {
// 	var payBody models.PayBody
// 	if err := c.BindJSON(&payBody); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	fromAddress := utils.CalculateHash(payBody.FromUser)
// 	toAddress := utils.CalculateHash(payBody.ToUser)
// 	currentBlock := Blockchain[len(Blockchain)-1]

// 	transaction := generateTransaction(fromAddress, toAddress, payBody.Price, currentBlock)

// 	spew.Dump(transaction)
// }

// func packageChain(person models.Person) models.Block {
// 	mutex.Lock()

// 	oldBlock := Blockchain[len(Blockchain)-1]

// 	hashes := services.GetAllTransactionsByBlock(transactions, oldBlock)

// 	newBlock := generateBlock(oldBlock, person, hashes)

// 	mutex.Unlock()

// 	if utils.IsBlockValid(newBlock, oldBlock) {
// 		candidateBlocks <- newBlock
// 	}

// 	spew.Dump(newBlock)

// 	return newBlock
// }

// // 生成区块
// func generateBlock(oldBlock models.Block, person models.Person, hashes []string) models.Block {

// 	var newBlock models.Block
// 	t := time.Now()

// 	address := utils.CalculateHash(person.Name)

// 	validators[address] = person.Balance

// 	newBlock.Timestamp = t.String()
// 	newBlock.Index = oldBlock.Index + 1
// 	newBlock.PreHash = oldBlock.Hash
// 	newBlock.Hash = utils.CalculateBlockHash(newBlock)
// 	newBlock.Validator = address
// 	newBlock.MerkleRoot = merkletree.CalculateMarkleRoot(hashes)

// 	return newBlock
// }

// func PickWinner() {
// 	lotteryPool := []string{}
// 	time.Sleep(3 * time.Second)
// 	var Winner string

// 	if len(tempBlocks) > 0 {
// 		for _, block := range tempBlocks {
// 			k, ok := validators[block.Validator]
// 			if ok {
// 				for i := 0; i < k; i++ {
// 					lotteryPool = append(lotteryPool, block.Validator)
// 				}
// 			}
// 		}

// 		Winner = lotteryPool[rand.Intn(len(lotteryPool))]

// 		for _, block := range tempBlocks {
// 			if block.Validator == Winner {
// 				mutex.Lock()
// 				Blockchain = append(Blockchain, block)
// 				mutex.Unlock()
// 			}
// 		}
// 	}

// 	fmt.Println("Winner " + Winner)

// 	tempBlocks = []models.Block{}
// 	tempTransactions = []models.Transaction{}
// }

// var transactions []models.Transaction
// var tempTransactions []models.Transaction

// func generateTransaction(fromAddress, toAddress string, price float64, block models.Block) models.Transaction {
// 	var transaction models.Transaction

// 	transaction.From = fromAddress
// 	transaction.To = toAddress
// 	transaction.GasPrice = strconv.FormatFloat(price, 'E', -1, 64)
// 	transaction.BlockHash = block.Hash
// 	transaction.Timestamp = time.Now().String()

// 	str := fromAddress + toAddress + transaction.Timestamp
// 	transaction.Hash = utils.CalculateHash(fromAddress + str)

// 	transactions = append(transactions, transaction)

// 	tempTransactions = append(tempTransactions, transaction)

// 	return transaction
// }
