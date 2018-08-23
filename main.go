package main

import (
	"block_chain/controllers"
	"block_chain/db"
	"block_chain/log"
	"block_chain/models"
	"block_chain/services"
	"block_chain/system"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var environment, port, configFilePath string

const (
	environmentDesc    = "Specifies the environment to run this server"
	portDesc           = "Runs Gin on the specified port"
	configFilePathDesc = "Uses a custom configuration"
)

func init() {
	flag.StringVar(&environment, "e", "development", environmentDesc)
	flag.StringVar(&port, "p", "5000", portDesc)
	flag.StringVar(&configFilePath, "c", "", configFilePathDesc)

	flag.Parse()

	if err := system.LoadConfiguration(configFilePath, environment); err != nil {
		log.Logger().Error(err.Error())
	}

	log.Logger().Infof("Starting server On %s ..", environment)
}

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Logger().Errorf("err open databases", err)
		return
	}
	defer db.Close()

	createTables(db)

	router := gin.Default()

	services.GenesisBlock()

	router.GET("/", controllers.HandleGetBlockChain)
	router.POST("/pay", controllers.HandlePay)

	go func() {
		for {
			services.PickWinner()
		}
	}()

	router.Run(":5000")
}

func createTables(db *gorm.DB) {
	// db.DropTableIfExists(&models.Transaction{}, &models.Block{})

	if !db.HasTable(&models.Block{}) {
		db.CreateTable(&models.Block{})
		log.Logger().Info("CreateTable Block")
	}

	if !db.HasTable(&models.Transaction{}) {
		db.CreateTable(&models.Transaction{})
		log.Logger().Info("CreateTable Transaction")
	}

	db.AutoMigrate(&models.Transaction{}, &models.Block{})
}

// 公司账户 1000万token
