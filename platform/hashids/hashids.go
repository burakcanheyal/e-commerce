package hashids

import (
	"github.com/speps/go-hashids"
)

type HashId struct {
	hd *hashids.HashID
}

var hId HashId

func init() {
	hd := hashids.NewData()
	hd.Salt = "Tahmin edilmesi çok güç bir salt"
	hd.MinLength = 30
	hid, _ := hashids.NewWithData(hd)
	hId = NewHashId(hid)
}

func NewHashId(hd *hashids.HashID) HashId {
	h := HashId{hd}
	return h
}
func EncodeId(id int) (string, error) {

	hashedIds, err := hId.hd.Encode([]int{id})
	if err != nil {
		return "", err
	}
	return hashedIds, nil
}
func DecodeId(id string) ([]int, error) {
	numbers, err := hId.hd.DecodeWithError(id)
	if err != nil {
		return nil, err
	}
	return numbers, err
}
