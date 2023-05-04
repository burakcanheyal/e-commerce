package handler

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/enum"
	"attempt4/internal/domain/service"
	"attempt4/platform/validation"
	"attempt4/platform/zap"
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
		zap.Logger.Error(err)
		context.JSON(http.StatusServiceUnavailable, ErrorInJson())
		return
	}

	userDto, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.FailInTokenParse)
		context.JSON(401, internal.UserNotFound)
		return
	}

	err := validation.ValidateStruct(wallet)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	err = w.walletService.UpdateBalance(wallet, userDto.Id)
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, NewHttpError(err))
		return
	}

	zap.Logger.Info("Cüzdan bakiye güncelleme isteği başarılı")
	context.JSON(http.StatusOK, SuccessInUpdate())
}

func (w *WalletServerHandler) CompletePurchase(context *gin.Context) {

	userDto, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	err := w.walletService.Purchase(userDto.Id)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	zap.Logger.Info("Ödeme tamamlama isteği başarılı")
	context.JSON(http.StatusOK, SuccessInPurchase())
}

func (w *WalletServerHandler) GetAllBuyTransactions(context *gin.Context) {
	userDto, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	items, total, err := w.walletService.GetAllTransactions(userDto.Id, enum.WalletBuyType)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}

	zap.Logger.Info("Tüm alım işlemlerini gösterme isteği başarılı")
	context.JSON(http.StatusOK, gin.H{"Toplam sipariş sayısı ": total, "Siparişler: ": items})
}

func (w *WalletServerHandler) GetAllSellTransactions(context *gin.Context) {
	user, exist := context.Keys["user"].(dto.TokenUserDto)
	if exist != true {
		zap.Logger.Error(internal.UserNotFound)
		context.JSON(401, internal.UserNotFound)
		return
	}

	pdf, err := w.walletService.ShowStatistics(user.Id)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewHttpError(err))
		return
	}
	context.Writer.Header().Set("Content-Type", "application/pdf")
	context.Writer.Header().Set("Content-disposition", "attachment; filename=kittens.pdf")

	zap.Logger.Info("Tüm satış işlemlerini gösterme isteği başarılı")
	context.Data(http.StatusOK, "application/octet-stream", pdf)
}
