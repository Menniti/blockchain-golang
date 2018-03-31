package model

import "sync"

//Difficulty the value that represents the number of 0 in front of hash to be mined
const Difficulty = 1

//Mutex prevent some go routine try to access some data at the same time, in this case we will use to prevent some block has been create at the same time
var Mutex = &sync.Mutex{}

//Block contains the data from block
type Block struct {
	//Is the position on datarecord on blockchain
	Index int
	//TimeStaping is the time when the data has been written
	TimeStamp string
	//BPM beats per minute (related to tutorial example) Medicine blockchain
	BPM int
	//Hash (sha256) cryptograph representation of actual record on the chain
	Hash string
	//PrevHash (sha256) cryptograph representation of previous record on the chain
	PrevHash string
	//Difficulty the dificult to mining block
	Difficulty int
	//Nonce the add value to achieve the hash
	Nonce string
}
