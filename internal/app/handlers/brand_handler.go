package handlers

import (
	dtos "crud/internal/app/dtos"
	services "crud/internal/app/services"
	utils "crud/internal/app/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	service services.BrandService
}

func NewBrandHandler(serv services.BrandService) *BrandHandler {
	return &BrandHandler{serv}
}

func (h *BrandHandler) CreateBrand(c *gin.Context) {
	var b dtos.CreateBrandDto
	err := c.BindJSON(&b)
	if err != nil {
		msg := fmt.Sprintf("Erro ao ler o corpo da requisição: %v", err)
		utils.JSON(c.Writer, http.StatusBadRequest, utils.Response{Error: msg})
		return
	}

	savedBrand, err := h.service.Create(&b)
	if err != nil {
		utils.JSON(c.Writer, http.StatusBadRequest, utils.Response{Error: err.Error()})
	}

	utils.JSON(c.Writer, http.StatusOK, utils.Response{Data: savedBrand})
}

func (h *BrandHandler) FindById(c *gin.Context) {
	brandId := c.Param("id")
	brand, _ := h.service.FindById(brandId)
	utils.JSON(c.Writer, http.StatusOK, utils.Response{Data: brand})
}

func (h *BrandHandler) FindAll(c *gin.Context) {
	response, err := h.service.FindAll()
	if err != nil {
		msg := fmt.Sprintf("Erro ao buscar marcas: %v", err)
		utils.JSON(c.Writer, http.StatusInternalServerError, utils.Response{Error: msg})
		return
	}

	utils.JSON(c.Writer, http.StatusOK, utils.Response{Data: response})
}

func (h *BrandHandler) UpdateById(c *gin.Context) {
	brandId := c.Param("id")
	var dto dtos.UpdateBrandDto
	err := c.BindJSON(&dto)
	if err != nil {
		msg := fmt.Sprintf("Erro ao ler o corpo da requisição: %v", err)
		utils.JSON(c.Writer, http.StatusBadRequest, utils.Response{Error: msg})
		return
	}

	response, err := h.service.UpdateById(brandId, &dto)
	if err != nil {
		msg := fmt.Sprintf("Erro ao atualizar a marca com ID %s: %v", brandId, err)
		utils.JSON(c.Writer, http.StatusBadRequest, utils.Response{Error: msg})
		return
	}

	utils.JSON(c.Writer, http.StatusOK, utils.Response{Data: response})
}

func (h *BrandHandler) DeleteById(c *gin.Context) {
	brandId := c.Param("id")
	if brandId == "" {
		utils.JSON(c.Writer, http.StatusBadRequest, utils.Response{Error: "ID não fornecido"})
		return
	}

	log.Println("brandId: ", brandId)
	err := h.service.DeleteById(brandId)
	if err != nil {
		msg := fmt.Sprintf("Erro ao deletar a marca com ID %s: %v", brandId, err)
		utils.JSON(c.Writer, http.StatusNotFound, utils.Response{Error: msg})
		return
	}

	utils.JSON(c.Writer, http.StatusOK, utils.Response{Data: "Marca deletada com sucesso"})
}
