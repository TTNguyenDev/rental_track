package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/TTNguyenDev/rental_track/db/sqlc"
)

type createRentalUnitRequest struct {
	HouseID int32  `json:"house_id" binding:"required"`
	Price   string `json:"price"    binding:"required"`
	Status  string `json:"status"   binding:"required,oneof=Rented Empty"`
}

func (server *Server) createRentalUnit(ctx *gin.Context) {
	var req createRentalUnitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateRentalUnitParams{
		HouseID: req.HouseID,
		Price:   req.Price,
		Status:  db.Rentalstatus(req.Status),
	}

	rentalUnit, err := server.store.CreateRentalUnit(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rentalUnit)
}

type getRentalUnitRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getRentalUnit(ctx *gin.Context) {
	var req getRentalUnitRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	rentalUnit, err := server.store.GetHouse(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, rentalUnit)
}

type getRentalUnitsByHouseRequest struct {
	HouseID  int32 `json:"house_id" binding:"required"`
	PageID   int32 `                binding:"required,min=1"        form:"page_id"`
	PageSize int32 `                binding:"required,min=5,max=10" form:"page_size"`
}

func (server *Server) getRentalUnitsByHouse(ctx *gin.Context) {
	var req getRentalUnitsByHouseRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetRentalUnitsByHouseParams{
		HouseID: req.HouseID,
		Limit:   req.PageSize,
		Offset:  (req.PageID - 1) * req.PageSize,
	}

	houses, err := server.store.GetRentalUnitsByHouse(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, houses)
}
