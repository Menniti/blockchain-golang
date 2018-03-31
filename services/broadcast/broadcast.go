package broadcast

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/menniti/blockchain-golang/model"
)

//ReceiveBroadcast funcao que recebe o broadcast
func ReceiveBroadcast(conn net.Conn, blockchain []model.Block, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		time.Sleep(30 * time.Second)
		model.Mutex.Lock()
		output, err := json.Marshal(blockchain)
		if err != nil {
			fmt.Println("[ReceiveBroadcast] Error to Marshal blockchain data", err.Error())
			log.Fatal(err)
		}
		model.Mutex.Unlock()
		io.WriteString(conn, string(output))

	}
}
