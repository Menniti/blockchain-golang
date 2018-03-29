package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/menniti/blockchain-golang/services/block"
	"github.com/menniti/blockchain-golang/webserver"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	go block.GenerateGenesisBlock()

	err = webserver.Run()
	if err != nil {
		fmt.Print("[main] Error to run the blockchain", err.Error())
		return
	}
}
