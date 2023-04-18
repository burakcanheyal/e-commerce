package handler

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/service"
	"attempt4/platform/validation"
	"attempt4/platform/zap"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderServerHandler struct {
	orderService service.OrderService
}

func NewOrderServerHandler(orderService service.OrderService) OrderServerHandler {
	o := OrderServerHandler{orderService}
	return o
}

func (o *OrderServerHandler) Create(context *gin.Context) {
	order := dto.OrderDto{}
	if err := context.BindJSON(&order); err != nil {
		zap.Logger.Error(err)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	id, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	err := validation.ValidateStruct(order)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	orderDescription, err := o.orderService.CreateOrder(order, id.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	zap.Logger.Info("Sipariş oluşturuldu")
	context.JSON(http.StatusOK, gin.H{"Kullanıcı ": orderDescription.Username, "Ürünler: ": orderDescription.Products})
}

func (o *OrderServerHandler) GetById(context *gin.Context) {
	id := context.Param("id")
	orderId, _ := strconv.ParseInt(id, 10, 64)
	orderIdInt32 := int32(orderId)

	order, err := o.orderService.GetOrderById(orderIdInt32)
	if err != nil {
		context.JSON(http.StatusNotFound, NonExistItem())
		return
	}

	zap.Logger.Info("Id ile sipariş isteme başarılı")
	context.JSON(http.StatusOK, order)
}

func (o *OrderServerHandler) Update(context *gin.Context) {
	order := dto.OrderDto{}
	if err := context.BindJSON(&order); err != nil {
		zap.Logger.Error(err)
		context.JSON(http.StatusServiceUnavailable, ErrorInJson())
		return
	}

	id, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	err := validation.ValidateStruct(order)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	err = o.orderService.UpdateOrder(order, id.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	zap.Logger.Info("Sipariş güncelleme başarılı")
	context.JSON(http.StatusOK, SuccessInUpdate())
}

func (o *OrderServerHandler) Delete(context *gin.Context) {
	order := dto.OrderDto{}
	if err := context.BindJSON(&order); err != nil {
		zap.Logger.Error(err)
		context.JSON(http.StatusServiceUnavailable, ErrorInJson())
		return
	}

	err := o.orderService.DeleteOrder(order.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	zap.Logger.Info("Siarpiş silme başarılı")
	context.JSON(http.StatusOK, SuccessInDelete())
}

func (o *OrderServerHandler) GetAllOrders(context *gin.Context) {
	filter := dto.Filter{}
	if err := context.ShouldBind(&filter); err != nil {
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	pagination := dto.Pagination{}
	if err := context.ShouldBind(&pagination); err != nil {
		zap.Logger.Error(err)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	id, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	orderDto, totalNumber, err := o.orderService.GetAllOrders(id.Id, filter, pagination)
	if err != nil {
		zap.Logger.Error(internal.FailInTokenParse)
		context.JSON(http.StatusNotFound, NonExistItem())
		return
	}

	zap.Logger.Info("Tüm siparişleri görüntüleme başarılı")
	context.JSON(http.StatusOK, gin.H{"Toplam Sipariş sayısı": totalNumber, "Siparişler: ": orderDto})
}
