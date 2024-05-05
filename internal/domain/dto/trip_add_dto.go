package dto

type TripAddDto struct {
	Id              int32   `json:"id"`
	Description     string  `json:"description"`
	Name            string  `json:"name"`
	Status          int32   `json:"status"`
	Lat             float64 `json:"lat"`
	Lng             float64 `json:"lng"`
	NaturePoint     int8    `json:"nature_point"`
	HistoricalPoint int8    `json:"historical_point"`
	AdventurePoint  int8    `json:"adventure_point"`
	CulturalPoint   int8    `json:"cultural_point"`
	RelaxPoint      int8    `json:"relax_point"`
}
