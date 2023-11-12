package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	db "github.com/TTNguyenDev/rental_track/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.Use(cors.Default())

	// Admins

	// House
	router.POST("/house", server.createHouse)
	router.GET("/house/:id", server.getHouse)
	router.GET("/houses", server.getHouses)

	// Rental Unit
	router.POST("/rentalunit", server.createRentalUnit)
	router.GET("/rentalunit/:id", server.getRentalUnit)
	router.GET("/rentalunitsByHouse", server.getRentalUnitsByHouse)

	// Renter
	router.POST("/renter", server.createRenter)
	router.GET("/renter/:id", server.getRenter)
	router.GET("/renters", server.getRenters)

	// Rental Agreement
	router.POST("/rentalAgreement", server.createRentalAgreement)
	router.GET("/rentalAgreement/:id", server.getRentalAgreement)
	router.GET("/rentalAgreements", server.getRentalAgreementsByRenter)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
