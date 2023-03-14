package validation

import (
	"errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validation struct {
	validator *validator.Validate
	uni       *ut.UniversalTranslator
}

func NewValidation(validator *validator.Validate, uni *ut.UniversalTranslator) Validation {
	v := Validation{validator, uni}
	return v
}
func (v *Validation) ValidateStruct(s interface{}) error {
	trans, _ := v.uni.GetTranslator("tr")
	//Todo:RegisterDefaultTranslations

	err := v.validator.Struct(s)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			return errors.New(e.Translate(trans))
		}
	}
	return nil
}
func (v *Validation) ValidatorCustomMessages() {
	trans, _ := v.uni.GetTranslator("tr")
	v.validator.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} bir değere sahip olmalıdır!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})
	v.validator.RegisterTranslation("lte", trans, func(ut ut.Translator) error {
		return ut.Add("lte", "{0} beklenen karakterden fazla giriş!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("lte", fe.Field())

		return t
	})
	v.validator.RegisterTranslation("gte", trans, func(ut ut.Translator) error {
		return ut.Add("gte", "{0} beklenen karakterden az giriş!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("gte", fe.Field())

		return t
	})
	v.validator.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} email formatına uygun olmayan giriş!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())

		return t
	})
}
