package handler

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/service"
	"attempt4/platform/zap"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminPanelHandler struct {
	adminService service.AdminService
}

func NewAdminPanelHandler(adminService service.AdminService) AdminPanelHandler {
	u := AdminPanelHandler{adminService}
	return u
}
func (a *AdminPanelHandler) GetUser(context *gin.Context) {
	user, err := a.adminService.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusNotFound, NewHttpError(err))
		return
	}

	zap.Logger.Info("Kullanıcı bilgileri görüntüleme başarılı")
	context.JSON(http.StatusOK, user)
}
func (a *AdminPanelHandler) GetTrip(context *gin.Context) {
	user, err := a.adminService.GetAllTrips()
	if err != nil {
		context.JSON(http.StatusNotFound, NewHttpError(err))
		return
	}

	zap.Logger.Info("Kullanıcı bilgileri görüntüleme başarılı")
	context.JSON(http.StatusOK, user)
}
func (a *AdminPanelHandler) DeleteUser(context *gin.Context) {
	code := dto.IdDto{}
	if err := context.BindJSON(&code); err != nil {
		zap.Logger.Error(internal.FailInTokenParse)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := a.adminService.DeleteUser(code.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	zap.Logger.Info("Kullanıcı silme Başarılı")
	context.JSON(http.StatusOK, SuccessInDelete())
}
func (a *AdminPanelHandler) DeleteTrip(context *gin.Context) {
	code := dto.IdDto{}
	if err := context.BindJSON(&code); err != nil {
		zap.Logger.Error(internal.FailInTokenParse)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := a.adminService.DeleteTrip(code.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	zap.Logger.Info("Trip silme Başarılı")
	context.JSON(http.StatusOK, SuccessInDelete())
}
func (a *AdminPanelHandler) CreateTrip(context *gin.Context) {
	trip := dto.TripAddDto{}
	if err := context.BindJSON(&trip); err != nil {
		zap.Logger.Error(err)
		context.JSON(http.StatusBadRequest, ErrorInJson())
		return
	}

	err := a.adminService.AddTrip(trip)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	zap.Logger.Info("Kullanıcı oluşturma başarılı")
	context.JSON(http.StatusOK, SuccessInCreate())
}
