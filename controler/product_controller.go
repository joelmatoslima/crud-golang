package controller

import (
	"fmt"
	"lsport/model"
	"lsport/usecase"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type productController struct {
	//usecase
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {

	//
	return productController{
		productUseCase: usecase,
	}

}

func (p productController) GetProducts(ctx *gin.Context) {

	var products, err = p.productUseCase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

	}

	ctx.JSON(http.StatusOK, products)

}

func (p productController) CreateProduct(ctx *gin.Context) {

	var product model.Product

	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

	}

	ctx.JSON(http.StatusCreated, insertedProduct)

}

func (p productController) GetProductById(ctx *gin.Context) {

	productId := ctx.Param("id")
	fmt.Println("productId:", productId)
	if strings.TrimSpace(productId) == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "O id precisa nao pode ser vazio",
		})
		return
	}
	// strconv
	var id, err = strconv.Atoi(productId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "I id do produto precisa ser um numero",
		})
		return

	}

	product, err := p.productUseCase.GetProductById(id)
	if err != nil {
		fmt.Println("Error GetProductById => ", err)
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Error do servidor",
		})
		return

	}
	fmt.Println("product:", product)

	if product == nil {

		fmt.Println("Error GetProductById => ", err)
		ctx.JSON(http.StatusNotFound, model.Response{
			Message: "Produto nao encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, product)

}
