package validation

import (
	"attempt4/platform/zap"
	"errors"
	"github.com/go-playground/locales/tr"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var v = validator.New()
var language = tr.New()
var uni = ut.New(language, language)

func init() {
	ValidatorCustomMessages()
}
func ValidateStruct(s interface{}) error {
	trans, _ := uni.GetTranslator("tr")

	err := v.Struct(s)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			tempErr := errors.New(e.Translate(trans))
			zap.Logger.Error(tempErr)
			return tempErr
		}
	}
	return nil
}
func ValidatorCustomMessages() {
	trans, _ := uni.GetTranslator("tr")
	v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} bir değere sahip olmalıdır!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})
	v.RegisterTranslation("lte", trans, func(ut ut.Translator) error {
		return ut.Add("lte", "{0} beklenen karakterden fazla giriş!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("lte", fe.Field())

		return t
	})
	v.RegisterTranslation("gte", trans, func(ut ut.Translator) error {
		return ut.Add("gte", "{0} beklenen karakterden az giriş!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("gte", fe.Field())

		return t
	})
	v.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} email formatına uygun olmayan giriş!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())

		return t
	})
}
