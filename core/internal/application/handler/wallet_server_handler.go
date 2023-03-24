package handler

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/service"
	"attempt4/core/platform/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WalletServerHandler struct {
	walletService service.WalletService
}

func NewWalletServerHandler(walletService service.WalletService) WalletServerHandler {
	w := WalletServerHandler{walletService}
	return w
}
func (w *WalletServerHandler) Update(context *gin.Context) {
	wallet := dto.WalletDto{}
	if err := context.BindJSON(&wallet); err != nil {
		context.JSON(http.StatusServiceUnavailable, ErrorInJson())
		return
	}

	id, exist := context.Keys["id"].(dto.IdDto)
	if exist != true {
		context.JSON(401, internal.UserNotFound)
		return
	}

	err := validation.ValidateStruct(wallet)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	err = w.walletService.UpdateBalance(wallet, id.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInUpdate())
}
func (w *WalletServerHandler) CompletePurchase(context *gin.Context) {

}
