package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/menniti/blockchain-golang/model"
	"github.com/menniti/blockchain-golang/server"
	"github.com/menniti/blockchain-golang/services/block"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server.bcServer = make(chan []model.Block)

	go block.GenerateGenesisBlock()

	err = webserver.Run()
	if err != nil {
		fmt.Print("[main] Error to run the blockchain", err.Error())
		return
	}
}
