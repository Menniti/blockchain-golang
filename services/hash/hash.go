package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/menniti/blockchain-golang/model"
)

//CalculateHash calcules the hash of block
func CalculateHash(block *model.Block) string {
	record := string(block.Index) + string(block.BPM) + block.PrevHash + block.TimeStamp + block.Nonce
	sha := sha256.New()
	sha.Write([]byte(record))
	hashed := sha.Sum(nil)
	return hex.EncodeToString(hashed)
}

//IsValidHash check if the hash has number of 0 in front of hash
func IsValidHash(hash string, difficult int) bool {
	prefix := strings.Repeat("0", model.Difficulty)
	fmt.Println("=========================")
	fmt.Println("prefix", prefix)
	fmt.Println("=========================")
	return strings.HasPrefix(hash, prefix)
}
