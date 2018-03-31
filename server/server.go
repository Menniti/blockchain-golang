package webserver

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/menniti/blockchain-golang/model"
	"github.com/menniti/blockchain-golang/router"
)

var bcServer chan []model.Block

//Run run the webservice
func Run() (err error) {

	mux := router.MakeMuxRouter()
	fmt.Print(os.Getenv("ADDR"))
	httpAddr := os.Getenv("ADDR")
	log.Println("Listen on: ", os.Getenv("ADDR"))
	server := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("[run] Error on the server", err.Error())
		return err
	}

	return nil
}
