package handler

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/service"
	"attempt4/platform/validation"
	"attempt4/platform/zap"
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
		zap.Logger.Error(err)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := validation.ValidateStruct(product)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	user, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	product, err = p.productService.CreateProduct(product, user.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, ItemNotAdded())
		return
	}

	zap.Logger.Info("Ürün oluşturma başarılı")
	context.JSON(http.StatusOK, product)
}

func (p *ProductServerHandler) GetByName(context *gin.Context) {
	name := context.Param("name")

	user, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	pro, err := p.productService.GetProductByName(name, user.Id)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	zap.Logger.Info("İsim ile ürün çağırma başarılı")
	context.JSON(http.StatusOK, pro)
}

func (p *ProductServerHandler) Update(context *gin.Context) {
	product := dto.ProductUpdateDto{}
	if err := context.BindJSON(&product); err != nil {
		zap.Logger.Error(err)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := validation.ValidateStruct(product)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	user, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	err = p.productService.UpdateProduct(product, user.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	zap.Logger.Info("Ürün güncelleme başarılı")
	context.JSON(http.StatusOK, SuccessInUpdate())
}

func (p *ProductServerHandler) Delete(context *gin.Context) {
	product := dto.ProductDto{}
	if err := context.BindJSON(&product); err != nil {
		zap.Logger.Error(err)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}
	user, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	err := p.productService.DeleteProduct(product.Name, user.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	zap.Logger.Info("Ürün silme başarılı")
	context.JSON(http.StatusOK, SuccessInDelete())
}

func (p *ProductServerHandler) GetAllProducts(context *gin.Context) {
	filter := dto.Filter{}
	if err := context.ShouldBind(&filter); err != nil {
		zap.Logger.Error(err)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	pagination := dto.Pagination{}
	if err := context.ShouldBind(&pagination); err != nil {
		zap.Logger.Error(err)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}
	user, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	products, totalNumber, err := p.productService.GetAllProducts(filter, pagination, user.Id)
	if err != nil {
		context.JSON(http.StatusNotFound, NonExistItem())
		return
	}

	zap.Logger.Info("Tüm ürünleri görüntüleme başarılı")
	context.JSON(http.StatusOK, gin.H{"Toplam ürün sayısı: ": totalNumber, "Ürünler:": products})
}
