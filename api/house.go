package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/TTNguyenDev/rental_track/db/sqlc"
)

type createHouseRequest struct {
	Name    string `json:"name"    binding:"required"`
	Address string `json:"address" binding:"required"`
	Kind    string `json:"kind"    binding:"required,oneof=House Rooms"`
}

func (server *Server) createHouse(ctx *gin.Context) {
	var req createHouseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateHouseParams{
		Name:    req.Name,
		Address: req.Address,
		Kind:    db.Housekind(req.Kind),
	}

	house, err := server.store.CreateHouse(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, house)
}

type getHouseRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getHouse(ctx *gin.Context) {
	var req getHouseRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	house, err := server.store.GetHouse(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, house)
}

type getHousesRequest struct {
	PageID   int32 `form:"page_id"   binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getHouses(ctx *gin.Context) {
	var req getHousesRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetHousesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	houses, err := server.store.GetHouses(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, houses)
}
