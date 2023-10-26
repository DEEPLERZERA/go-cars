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

type getCarRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getCar(ctx *gin.Context) {
	var req getCarRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	car, err := server.store.GetCar(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, car)
}

type deleteCarRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteCar(ctx *gin.Context) {
	var req deleteCarRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.DeleteCar(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, true)
}

type updateCarRequest struct {
	ID    int32  `json:"id" binding:"required"`
	Name  string `json:"name"`
	Price int32  `json:"price"`
	Brand string `json:"brand"`
}

func (server *Server) updateCar(ctx *gin.Context) {
	var req updateCarRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCarParams{
		ID:    req.ID,
		Name:  req.Name,
		Price: req.Price,
		Brand: req.Brand,
	}

	car, err := server.store.UpdateCar(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, car)
}

func (server *Server) getCars(ctx *gin.Context) {
	cars, err := server.store.GetCars(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, cars)
}