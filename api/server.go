package api

import (
	"fmt"

	"github.com/akhil-is-watching/goshort/storage"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	listenAddr string
	router     *gin.Engine
	store      storage.Storage
}

func NewAPIServer(listenAddr string, store storage.Storage) *APIServer {

	router := gin.Default()

	return &APIServer{
		listenAddr: listenAddr,
		router:     router,
		store:      store,
	}
}

func (s *APIServer) Run() {
	s.router.Run(s.listenAddr)
	fmt.Println("Running on: ", s.listenAddr)
}
