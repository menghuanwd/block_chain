package models

import (
	"block_chain/db"
	"github.com/lib/pq"
)

type Transaction struct {
	Timestamp string `json:"timestamp"`
	Hash      string `json:"hash" gorm:"type:varchar(100);unique_index"`
	From      string `json:"from"`
	To        string `json:"to"`
	GasPrice  string `json:"gas_price"`
	Nonce     string `json:"nonce"`
	BlockHash string `json:"block_hash"`
}

type PayBody struct {
	FromUser string  `json:"from_user"`
	ToUser   string  `json:"to_user"`
	Price    float64 `json:"price"`
}

type Block struct {
	Index        uint           `json:"index"`
	Timestamp    string         `json:"timestamp"`
	Hash         string         `json:"hash" gorm:"type:varchar(100);unique_index"`
	MerkleRoot   string         `json:"merkle_root"`
	Transactions pq.StringArray `json:"transactions" gorm:"type:varchar(100)[]"`
	PreHash      string         `json:"pre_hash"`
	Validator    string         `json:"validator"`
}

// Create ...
func (block *Block) Create() (*Block, error) {
	err := db.GetDB().Create(&block).Error

	return block, err
}

// Create ...
func (transaction *Transaction) Create() (*Transaction, error) {
	err := db.GetDB().Create(&transaction).Error

	return transaction, err
}
