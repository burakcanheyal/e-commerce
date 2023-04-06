package handler

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/enum"
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

	id, exist := context.Keys["user"].(dto.TokenUserDto)
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
	id, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		context.JSON(401, internal.UserNotFound)
		return
	}

	err := w.walletService.Purchase(id.Id)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInPurchase())

}

func (w *WalletServerHandler) GetAllBuyTransactions(context *gin.Context) {
	id, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		context.JSON(401, internal.UserNotFound)
		return
	}

	items, total, err := w.walletService.GetAllTransactions(id.Id, enum.WalletBuyType)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, gin.H{"Toplam sipariş sayısı ": total, "Siparişler: ": items})

}

func (w *WalletServerHandler) GetAllSellTransactions(context *gin.Context) {
	user, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		context.JSON(401, internal.UserNotFound)
		return
	}

	err := w.walletService.ShowStatistics(user.Id)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	context.JSON(http.StatusOK, SuccessInCreatingPdf())
}
