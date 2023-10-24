package api 

import (
	db "github.com/DEEPLERZERA/go-cars/db/sqlc"
	"github.com/gin-gonic/gin"
)


type Server struct {
	store *db.ExecuteStore 
	router *gin.Engine
}

func InstanceServer(store *db.ExecuteStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"Api has one error:": err.Error()}
}