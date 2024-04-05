package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateStructSuccess(t *testing.T) {
	test := struct {
		TestString string  `validate:"required,gte=1,lte=32"`
		TestInt    int32   `validate:"required,lte=127,gte=1,number"`
		TestDouble float32 `validate:"required,gte=1,number"`
	}{
		TestString: "TestString",
		TestInt:    12,
		TestDouble: 11.1,
	}
	err := ValidateStruct(test)
	assert.Equal(t, err, nil)
}
func TestValidateStructGteFail(t *testing.T) {
	test := struct {
		TestString string  `validate:"required,gte=1,lte=32"`
		TestInt    int32   `validate:"required,lte=127,gte=5,number"`
		TestDouble float32 `validate:"required,gte=1,number"`
	}{
		TestString: "TestString",
		TestInt:    4,
		TestDouble: 11.1,
	}
	err := ValidateStruct(test)
	assert.NotEqual(t, err, nil)
}
func TestValidateStructLteFail(t *testing.T) {
	test := struct {
		TestDouble float32 `validate:"required,gte=1,lte=10,number"`
	}{
		TestDouble: 11.1,
	}
	err := ValidateStruct(test)
	assert.NotEqual(t, err, nil)
}
func TestValidateStructRequiredFail(t *testing.T) {
	test := struct {
		TestDouble float32 `validate:"required,gte=1,number"`
	}{}
	err := ValidateStruct(test)
	assert.NotEqual(t, err, nil)
}
func TestValidateStructEmailFail(t *testing.T) {
	test := struct {
		TestString string `validate:"required,email"`
	}{
		TestString: "TestString",
	}
	err := ValidateStruct(test)
	assert.NotEqual(t, err, nil)
}
