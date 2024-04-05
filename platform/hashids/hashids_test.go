package hashids

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeIdSuccess(t *testing.T) {
	_, err := EncodeId(123)
	assert.Equal(t, err, nil)
}
func TestDecodeIdSuccess(t *testing.T) {
	encodedId, _ := EncodeId(123)
	decodedId, err := DecodeId(encodedId)

	assert.Equal(t, err, nil)
	assert.Equal(t, decodedId[0], 123)

}

func TestDecodeIdError(t *testing.T) {
	encodedId, _ := EncodeId(123)
	decodedId, _ := DecodeId(encodedId)

	assert.NotEqual(t, decodedId[0], 456)
}
