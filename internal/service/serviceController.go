package service

import (
	"github.com/ErayOzdayioglu/api-gateway/internal/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServiceController struct {
	Repository ServiceRepository
}

func (c *ServiceController) CreateService(ctx *gin.Context) {
	var request AddServiceRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error()))
		return
	}
	result, err := c.Repository.Save(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, common.SuccessResponse(result))
}

func (c *ServiceController) FindByName(ctx *gin.Context) {
	name := ctx.Param("name")

	result, err := c.Repository.FindByName(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.SuccessResponse(result))
}
