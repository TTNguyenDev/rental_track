package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	db "github.com/TTNguyenDev/rental_track/db/sqlc"
)

type createRentalAgreementRequest struct {
	RenterID  int32     `json:"renter_id"  binding:"required"`
	RentalID  int32     `json:"rental_id"  binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date"   binding:"required"`
	Price     string    `json:"price"      binding:"required"`
}

func (server *Server) createRentalAgreement(ctx *gin.Context) {
	var req createRentalAgreementRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateRentalAgreementParams{
		RenterID:  req.RenterID,
		RentalID:  req.RentalID,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Price:     req.Price,
	}

	result, err := server.store.CreateRentalAgreement(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

type getRentalAgreementRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getRentalAgreement(ctx *gin.Context) {
	var req getRentalAgreementRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := server.store.GetRentalAgreement(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

type getRentalAgreementsByRenterRequest struct {
	RenterID int32 `json:"renter_id" binding:"required"`
	PageID   int32 `                 binding:"required,min=1"        form:"page_id"`
	PageSize int32 `                 binding:"required,min=5,max=10" form:"page_size"`
}

func (server *Server) getRentalAgreementsByRenter(ctx *gin.Context) {
	var req getRentalAgreementsByRenterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetRentalAgreementsByRenterParams{
		RenterID: req.RenterID,
		Limit:    req.PageSize,
		Offset:   (req.PageID - 1) * req.PageSize,
	}

	result, err := server.store.GetRentalAgreementsByRenter(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
