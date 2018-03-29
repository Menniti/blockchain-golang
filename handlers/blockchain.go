package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/davecgh/go-spew/spew"

	"github.com/menniti/blockchain-golang/model"
	"github.com/menniti/blockchain-golang/services/block"
	"github.com/menniti/blockchain-golang/services/chain"
)

//HandlerGetBlockchain return blockchain information
func HandlerGetBlockchain(w http.ResponseWriter, r *http.Request) {
	fmt.Print(chain.Blockchain)
	bytes, err := json.MarshalIndent(chain.Blockchain, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

//HandlerWriteBlock post new block on the chain
func HandlerWriteBlock(w http.ResponseWriter, r *http.Request) {
	message := model.Message{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&message); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	newBlock, err := block.GenerateBlock(chain.Blockchain[len(chain.Blockchain)-1], message.BMP)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, message)
		return
	}
	fmt.Println("===========================")
	fmt.Println("newblock: ", newBlock)
	fmt.Println("===========================")
	fmt.Println("chain.Blockchain: ", chain.Blockchain)
	fmt.Println("===========================")
	fmt.Println("len(chain.Blockchain): ", len(chain.Blockchain))
	fmt.Println("===========================")
	fmt.Println("chain.Blockchain[len(block.Blockchain)-1]: ", chain.Blockchain[len(chain.Blockchain)-1])
	fmt.Println("===========================")
	fmt.Println("IS VALID BLOCK ? ", block.IsValidBlock(newBlock, chain.Blockchain[len(chain.Blockchain)-1]))

	if block.IsValidBlock(chain.Blockchain[len(chain.Blockchain)-1], newBlock) {
		fmt.Println("IS VALID BLOCK ? ", block.IsValidBlock(newBlock, chain.Blockchain[len(chain.Blockchain)-1]))
		fmt.Println("===========================")
		chainOfBlocks := append(chain.Blockchain, newBlock)
		chain.ReplaceChain(chainOfBlocks, chain.Blockchain)
		spew.Dump(chain.Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func respondWithJSON(writter http.ResponseWriter, request *http.Request, statusCode int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		writter.WriteHeader(http.StatusInternalServerError)
		writter.Write([]byte("HTTP 500: INTERNAL SERVER ERROR"))
		return
	}
	writter.WriteHeader(statusCode)
	writter.Write(response)
}
