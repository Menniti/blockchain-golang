package block

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/menniti/blockchain-golang/model"
	"github.com/menniti/blockchain-golang/services/chain"
	"github.com/menniti/blockchain-golang/services/hash"
)

//GenerateBlock funcion that generates block
func GenerateBlock(oldBlock model.Block, BPM int) (newBlock model.Block, err error) {

	time := time.Now()

	newBlock.TimeStamp = time.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Index = oldBlock.Index + 1
	newBlock.Hash = hash.CalculateHash(&newBlock)
	newBlock.Difficulty = model.Difficulty

	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		fmt.Println("====================")
		fmt.Println("hex ", hex)
		fmt.Println("====================")
		newBlock.Nonce = hex

		fmt.Println("==================")
		fmt.Println("newBlock.Nonce", newBlock.Nonce)
		fmt.Println("==================")
		if !hash.IsValidHash(hash.CalculateHash(&newBlock), newBlock.Difficulty) {
			fmt.Println(hash.CalculateHash(&newBlock), " do more work")
			continue
		} else {
			fmt.Println(hash.CalculateHash(&newBlock), " work done")
			newBlock.Hash = hash.CalculateHash(&newBlock)
			break
		}
	}
	return newBlock, nil
}

//IsValidBlock check if the block is valid
func IsValidBlock(oldBlock, newBlock model.Block) bool {
	fmt.Println("oldBlock", oldBlock)
	fmt.Println("===========================")
	fmt.Println("newBlock", newBlock)
	fmt.Println("===========================")
	fmt.Print("calculateHash(&newBlock) != newBlock.Hash: ", hash.CalculateHash(&newBlock) != newBlock.Hash)
	fmt.Println("===========================")
	if hash.CalculateHash(&newBlock) != newBlock.Hash {
		fmt.Println("calculateHash Retorno falso")
		return false
	}
	fmt.Println("oldBlock.Hash != newBlock.PrevHash", oldBlock.Hash != newBlock.PrevHash)
	fmt.Println("===========================")
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	fmt.Println("oldBlock.Index+1 != newBlock.Index", oldBlock.Index+1 != newBlock.Index)
	fmt.Println("===========================")
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	return true
}

//GenerateGenesisBlock generate the genesisblock
func GenerateGenesisBlock() {
	t := time.Now()
	genesisBlock := model.Block{0, t.String(), 0, "", "", 0, ""}
	spew.Dump(genesisBlock)
	model.Mutex.Lock()
	chain.Blockchain = append(chain.Blockchain, genesisBlock)
	model.Mutex.Unlock()
}
