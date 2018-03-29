package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/menniti/blockchain-golang/handlers"
)

//MakeMuxRouter build a index router
func MakeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handlers.HandlerGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handlers.HandlerWriteBlock).Methods("POST")
	return muxRouter
}
