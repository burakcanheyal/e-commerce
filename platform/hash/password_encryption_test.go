package hash

import (
	"attempt4/internal"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestEncryptPasswordSuccess(t *testing.T) {
	_, err := EncryptPassword("TestData")
	assert.Equal(t, err, nil)
}

func TestCompareEncryptedPasswordsSuccess(t *testing.T) {
	testData, _ := EncryptPassword("TestData")

	err := CompareEncryptedPasswords(testData, "TestData")
	assert.Equal(t, err, nil)
}
func TestCompareEncryptedPasswordsInvalidPassword(t *testing.T) {
	testData, _ := EncryptPassword("TestData")

	err := CompareEncryptedPasswords(testData, "Invalid Password")
	assert.Equal(t, err, internal.InvalidPassword)
}
