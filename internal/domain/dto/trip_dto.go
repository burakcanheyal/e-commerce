package dto

type TripDto struct {
	Id          int32   `json:"id"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Status      int32   `json:"status"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
}
