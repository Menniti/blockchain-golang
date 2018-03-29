package chain

import "github.com/menniti/blockchain-golang/model"

//Blockchain contains the blockchain array
var Blockchain = []model.Block{}

//ReplaceChain replace the chain
func ReplaceChain(newBlocks, blockchain []model.Block) {

	if len(newBlocks) > len(blockchain) {
		Blockchain = newBlocks
	}
}
