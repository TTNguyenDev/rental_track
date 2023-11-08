package api

import (
	"github.com/gin-gonic/gin"

	db "github.com/TTNguyenDev/rental_track/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// TODO: add routes to router
	router.POST("/house", server.createHouse)
	router.GET("/house/:id", server.getHouse)
	router.GET("/houses", server.getHouses)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
