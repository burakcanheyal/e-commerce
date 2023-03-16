package hashids

import (
	"github.com/speps/go-hashids"
)

type HashId struct {
	hd *hashids.HashID
}

func NewHashId(hd *hashids.HashID) HashId {
	h := HashId{hd}
	return h
}
func (h *HashId) EncodeId(id int) (string, error) {

	hashedIds, err := h.hd.Encode([]int{id})
	if err != nil {
		return "", err
	}
	return hashedIds, nil
}
func (h *HashId) DecodeId(id string) ([]int, error) {
	numbers, err := h.hd.DecodeWithError(id)
	if err != nil {
		return nil, err
	}
	return numbers, err
}
