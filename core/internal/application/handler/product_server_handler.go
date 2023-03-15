package handler

import (
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/service"
	"attempt4/core/platform/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductServerHandler struct {
	productService service.ProductService
}

func NewProductServerHandler(productService service.ProductService) ProductServerHandler {
	p := ProductServerHandler{productService}
	return p
}

func (p *ProductServerHandler) Create(context *gin.Context) {
	product := dto.ProductDto{}
	if err := context.BindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := validation.ValidateStruct(product)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	tokenString := context.GetHeader("Authentication")
	if tokenString == "" {
		context.JSON(401, TokenError())
		return
	}

	product, err = p.productService.CreateProduct(product, tokenString)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, ItemNotAdded())
		return
	}

	context.JSON(http.StatusOK, product)
}
func (p *ProductServerHandler) GetByName(context *gin.Context) {
	name := context.Param("name")
	pro, err := p.productService.GetProductByName(name)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, pro)
}

func (p *ProductServerHandler) Update(context *gin.Context) {
	product := dto.ProductUpdateDto{}
	if err := context.BindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := validation.ValidateStruct(product)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	err = p.productService.UpdateProduct(product)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInUpdate())
}
func (p *ProductServerHandler) Delete(context *gin.Context) {
	product := dto.ProductDto{}
	if err := context.BindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := p.productService.DeleteProduct(product.Name)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInDelete())
}
func (p *ProductServerHandler) GetAllProducts(context *gin.Context) {
	filter := dto.Filter{}
	if err := context.ShouldBind(&filter); err != nil {
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	pagination := dto.Pagination{}
	if err := context.ShouldBind(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	products, totalNumber, err := p.productService.GetAllProducts(filter, pagination)
	if err != nil {
		context.JSON(http.StatusNotFound, NonExistItem())
		return
	}
	context.JSON(http.StatusOK, gin.H{"Toplam ürün sayısı: ": totalNumber, "Ürünler:": products})
}
