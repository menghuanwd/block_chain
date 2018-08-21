package services

import (
	"block_chain/db"
	"block_chain/merkletree"
	"block_chain/models"
	"block_chain/utils"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"math/rand"
	"strconv"
	"time"
)

// PickWinner ...
func PickWinner() {
	time.Sleep(5 * time.Second)
	fmt.Println("PickWinner now ....")

	validator := getWinnerAddress()

	if validator == "" {
		return
	}

	fmt.Println("PickWinner is ", validator)

	reward(validator)

	block := generateBlock(validator)

	addBlockHashToTransction(block)
}

// 奖励
func reward(validator string) {
	var paybody models.PayBody

	paybody.FromUser = "God"
	paybody.ToUser = validator
	paybody.Price = 20

	fromAddress := utils.CalculateHash(paybody.FromUser)

	CreateTransaction(fromAddress, validator, paybody.Price)
}

func getWinnerAddress() string {
	lotteryPool := []string{}

	transactions := GetUnblockTransactions()

	if len(transactions) == 0 {
		return ""
	}

	for _, transaction := range transactions {
		priceFloat64, _ := strconv.ParseFloat(transaction.GasPrice, 64)

		times := int(priceFloat64)

		for i := 0; i < times; i++ {
			lotteryPool = append(lotteryPool, transaction.From)
		}
	}

	winnerAddress := lotteryPool[rand.Intn(len(lotteryPool))]

	return winnerAddress
}

func award() {

}

// GenesisBlock ...
func GenesisBlock() models.Block {
	var block models.Block

	block.MerkleRoot = merkletree.CalculateMarkleRoot([]string{""})
	block.Transactions = nil
	block.PreHash = ""
	block.Validator = ""
	block.Timestamp = time.Now().String()
	block.Hash = calculateBlockHash(block)

	block.Create()

	return block
}

func generateBlock(validator string) models.Block {
	lastBlock := getLastBlock()

	hashes := GetUnblockTransactionsHashes()

	var block models.Block

	block.Index = lastBlock.Index + 1
	block.PreHash = lastBlock.Hash
	block.Timestamp = time.Now().String()
	block.Hash = calculateBlockHash(block)
	block.Validator = validator
	block.MerkleRoot = merkletree.CalculateMarkleRoot(hashes)
	block.Transactions = hashes

	block.Create()

	return block
}

func getLastBlock() models.Block {
	var lastBlock models.Block

	db.GetDB().Order("index desc").First(&lastBlock)

	return lastBlock
}

// GetUnblockTransactionsHashes 获取交易hashes
func GetUnblockTransactionsHashes() []string {
	var hashes []string

	transactions := GetUnblockTransactions()

	for _, transaction := range transactions {
		if transaction.BlockHash == "" {
			hashes = append(hashes, transaction.Hash)
		}
	}

	return hashes
}

// GetUnblockTransactions 获取未打包的交易
func GetUnblockTransactions() []models.Transaction {
	var transactions []models.Transaction

	db.GetDB().Where("block_hash = ?", "").Find(&transactions)

	return transactions
}

func addBlockHashToTransction(block models.Block) {
	var transactions []models.Transaction

	db.GetDB().Where("block_hash = ?", "").Find(&transactions).UpdateColumn("block_hash", block.Hash)
}

// CreateTransaction ...
func CreateTransaction(fromAddress, toAddress string, price float64) models.Transaction {
	var transaction models.Transaction

	nonce := utils.Nonce(16)

	fmt.Println(fromAddress, toAddress)

	transaction.From = fromAddress
	transaction.To = toAddress
	transaction.GasPrice = strconv.FormatFloat(price, 'E', -1, 64)
	transaction.Timestamp = time.Now().String()
	transaction.Nonce = nonce
	transaction.Hash = utils.CalculateHash(fromAddress + toAddress + nonce)

	spew.Dump(transaction)

	transaction.Create()

	return transaction
}

// calculateBlockHash ...
func calculateBlockHash(block models.Block) string {
	record := string(block.Index) + block.Timestamp + block.PreHash

	return utils.CalculateHash(record)
}

func IsBlockValid(block, lastBlock models.Block) bool {
	if lastBlock.Index+1 != block.Index {
		return false
	}

	if lastBlock.Hash != block.PreHash {
		return false
	}

	if calculateBlockHash(block) != block.Hash {
		return false
	}

	return true
}
