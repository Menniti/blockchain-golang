package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/menniti/blockchain-golang/handlers"
	"github.com/menniti/blockchain-golang/model"
	"github.com/menniti/blockchain-golang/services/block"
)

//BcServer get the chain
var BcServer chan []model.Block

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	BcServer = make(chan []model.Block)

	go block.GenerateGenesisBlock()

	//Listing direcly in the port ADDR
	server, err := net.Listen("tcp", ":"+os.Getenv("ADDR"))
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handlers.HandleConnection(conn, BcServer)
	}
}
