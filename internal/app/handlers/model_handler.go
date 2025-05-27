package handlers

import (
	dtos "crud/internal/app/dtos"
	services "crud/internal/app/services"
	utils "crud/internal/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ModelHandler struct {
	modelService services.ModelService
}

func NewModelHandler(modelServ services.ModelService) *ModelHandler {
	return &ModelHandler{
		modelService: modelServ,
	}
}

func (h *ModelHandler) CreateModel(c *gin.Context) {
	var createModelDto dtos.CreateModelDto
	if err := c.ShouldBindJSON(&createModelDto); err != nil {
		utils.JSON(c.Writer, http.StatusBadRequest, utils.Response{Error: err.Error()})
		return
	}
	result, _ := h.modelService.Create(&createModelDto)
	utils.JSON(c.Writer, http.StatusCreated, utils.Response{Data: result})
}

func (h *ModelHandler) FindAll(c *gin.Context) {
	result, _ := h.modelService.FindAll()
	utils.JSON(c.Writer, http.StatusOK, utils.Response{Data: result})
}

func (h *ModelHandler) FindById(c *gin.Context) {
	id := c.Param("id")
	result, err := h.modelService.FindById(id)
	if err != nil {
		utils.JSON(c.Writer, http.StatusNotFound, utils.Response{Error: err.Error()})
		return
	}
	utils.JSON(c.Writer, http.StatusOK, utils.Response{Data: result})
}

func (h *ModelHandler) UpdateById(c *gin.Context) {
	id := c.Param("id")
	var updateModelDto dtos.UpdateModelDto
	if err := c.ShouldBindJSON(&updateModelDto); err != nil {
		utils.JSON(c.Writer, http.StatusBadRequest, utils.Response{Error: err.Error()})
		return
	}
	result, err := h.modelService.UpdateById(id, &updateModelDto)
	if err != nil {
		utils.JSON(c.Writer, http.StatusNotFound, utils.Response{Error: err.Error()})
		return
	}
	utils.JSON(c.Writer, http.StatusOK, utils.Response{Data: result})
}

func (h *ModelHandler) DeleteById(c *gin.Context) {
	id := c.Param("id")
	err := h.modelService.DeleteById(id)
	if err != nil {
		utils.JSON(c.Writer, http.StatusNotFound, utils.Response{Error: err.Error()})
		return
	}
	utils.JSON(c.Writer, http.StatusNoContent, utils.Response{})
}
