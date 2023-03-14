package handler

type ApplicationError struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

func NonExistItem() ApplicationError {
	return ApplicationError{Result: "kayıt bulunamadı", Message: "Seçili kayıt bulunamadı"}
}
func ItemNotAdded() ApplicationError {
	return ApplicationError{Result: "Hata", Message: "Kayıt eklenemedi"}
}
func ErrorInJson() ApplicationError {
	return ApplicationError{Result: "Hata", Message: "Sistem Arızası"}
}
func TokenError() ApplicationError {
	return ApplicationError{Result: "Hata", Message: "Token Hatası"}
}
func UserExist() ApplicationError {
	return ApplicationError{Result: "Hata", Message: "Kayıtlı Kullanıcı"}
}
func SuccessInDelete() ApplicationError {
	return ApplicationError{Result: "Başarılı", Message: "Kayıt başarıyla silindi"}
}
func SuccessInUpdate() ApplicationError {
	return ApplicationError{Result: "Başarılı", Message: "Kayıt başarıyla güncellendi"}
}
func SuccessInCreate() ApplicationError {
	return ApplicationError{Result: "Başarılı", Message: "Kayıt başarıyla eklendi"}
}
func SuccessInActivate() ApplicationError {
	return ApplicationError{Result: "Başarılı", Message: "Kullanıcı aktif edildi"}
}
func NewHttpError(err error) ApplicationError {
	return ApplicationError{Result: "Hata", Message: err.Error()}
}
