package dto

type PieItemDto struct {
	Ratio           string `json:"Ratio"`
	OperationNumber string `json:"OperationNumber"`
}
type PieChartData struct {
	PieChartData []PieItemDto
}
