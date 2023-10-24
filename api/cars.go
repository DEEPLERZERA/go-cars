package api

import (
	"net/http"

	db "github.com/DEEPLERZERA/go-cars/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createCarRequest struct {
	Name  string `json:"name" binding:"required"`
	Price int32  `json:"price" binding:"required"`
	Brand string `json:"brand" binding:"required"`
}

func (server *Server) createCar(ctx *gin.Context) {
	var req createCarRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCarParams{
		Name:  req.Name,
		Price: req.Price,
		Brand: req.Brand,
	}

	car, err := server.store.CreateCar(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, car)
}
