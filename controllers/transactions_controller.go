package controllers

import (
	"block_chain/db"
	"block_chain/models"
	"block_chain/services"
	"block_chain/utils"
	// "github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var mutex = &sync.Mutex{}

// HandleGetBlockChain ...
func HandleGetBlockChain(c *gin.Context) {
	var blocks []models.Block

	db.GetDB().Find(&blocks)

	c.JSON(200, blocks)
}

// HandlePay ...
func HandlePay(c *gin.Context) {
	var payBody models.PayBody

	if err := c.BindJSON(&payBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// spew.Dump(payBody)

	fromAddress := utils.CalculateHash(payBody.FromUser)
	toAddress := utils.CalculateHash(payBody.ToUser)

	transaction := services.CreateTransaction(fromAddress, toAddress, payBody.Price)

	c.JSON(http.StatusCreated, transaction)
}
