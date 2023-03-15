package internal

import "errors"

var (
	DBNotCreated = errors.New("Kayıt oluşturulamadı")
	DBNotFound   = errors.New("Kayıt bulunamadı")
	DBNotUpdated = errors.New("Kayıt güncellenemedi")
	DBNotDeleted = errors.New("Kayıt silinemedi")

	ExceedOrder = errors.New("Varolandan fazla sayıda sipariş")

	FailInTokenParse = errors.New("Token parçalanamadı")
	FailInToken      = errors.New("Hatalı token")

	FailInHash = errors.New("Şifre Hash'lenemedi")

	InvalidPassword = errors.New("Hatalı şifre")

	UserNotCreated   = errors.New("Kayıtlı kullanıcı")
	UserExist        = errors.New("Kayıtlı kullanıcı")
	UserNotFound     = errors.New("Kullanıcı bulunamadı")
	DeletedUser      = errors.New("Silinmiş kullanıcı")
	PassiveUser      = errors.New("Pasif kullanıcı")
	UserUnactivated  = errors.New("Aktive edilmemiş kullanıcı")
	UserUnauthorized = errors.New("Bunun işlem için yetkili değilsiniz")

	FailInVerify     = errors.New("Aktivasyon kodu yanlış")
	ExceedVerifyCode = errors.New("Aktivasyon kodunun süresi dolmuş. Yeni kod gönderildi")

	ProductNotFound    = errors.New("Ürün bulunamadı")
	ProductExist       = errors.New("Kayıtlı ürün")
	ProductDeleted     = errors.New("Silinmiş ürün")
	ProductUnavailable = errors.New("Stokta kalmayan ürün")

	OrderNotFound = errors.New("Sipariş bulunamadı")
	OrderExist    = errors.New("Kayıtlı sipariş")

	KeyNotCreated = errors.New("Rol yaratılamadı")
)
